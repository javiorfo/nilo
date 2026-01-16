package nilo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Property string
}

func (t testStruct) Default() testStruct {
	return testStruct{"Default"}
}

func TestOption(t *testing.T) {
	t.Run("AsValue", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			assert.Equal(t, 42, opt.AsValue())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.Panics(t, func() {
				opt.AsValue()
			})
		})
	})

	t.Run("Or", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			assert.Equal(t, 42, opt.Or(24))
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.Equal(t, 24, opt.Or(24))
		})
	})

	t.Run("OrDefault", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(testStruct{"ok"})
			assert.Equal(t, "ok", opt.OrDefault().Property)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[testStruct]()
			assert.Equal(t, "Default", opt.OrDefault().Property)
		})
	})

	t.Run("OrError", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			value, err := opt.OrError(func() error { return errors.New("error") })
			assert.Equal(t, 42, *value)
			assert.NoError(t, err)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			value, err := opt.OrError(func() error { return errors.New("error") })
			assert.Error(t, err)
			assert.Nil(t, value)
		})
	})

	t.Run("Filter", func(t *testing.T) {
		t.Run("when value satisfies the filter", func(t *testing.T) {
			opt := Value(42)
			assert.Equal(t, 42, opt.Filter(func(i int) bool {
				return i > 0
			}).AsValue())
		})

		t.Run("when value does not satisfy the filter", func(t *testing.T) {
			opt := Value(42)
			assert.True(t, opt.Filter(func(i int) bool {
				return i < 0
			}).IsNil())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.True(t, opt.Filter(func(i int) bool {
				return i > 0
			}).IsNil())
		})
	})

	t.Run("IsNil", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			assert.False(t, opt.IsNil())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.True(t, opt.IsNil())
		})
	})

	t.Run("IsValue", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			assert.True(t, opt.IsValue())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.False(t, opt.IsValue())
		})
	})

	t.Run("Inspect", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			var result int
			opt.Inspect(func(i int) {
				result = i
			})
			assert.Equal(t, 42, result)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			var result int
			opt.Inspect(func(i int) {
				result = i
			})
			assert.Zero(t, result)
		})
	})

	t.Run("Consume", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			var result int
			opt.Consume(func(i int) {
				result = i
			})
			assert.Equal(t, 42, result)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			var result int
			opt.Consume(func(i int) {
				result = i
			})
			assert.Zero(t, result)
		})
	})

	t.Run("OrElse", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			assert.Equal(t, 42, opt.OrElse(func() int {
				return 24
			}))
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.Equal(t, 24, opt.OrElse(func() int {
				return 24
			}))
		})
	})

	t.Run("Nil", func(t *testing.T) {
		opt := Nil[int]()
		assert.True(t, opt.IsNil())
	})

	t.Run("Value", func(t *testing.T) {
		opt := Value(42)
		assert.Equal(t, 42, opt.AsValue())
	})

	t.Run("Nilable", func(t *testing.T) {
		t.Run("when value is not nil", func(t *testing.T) {
			value := 42
			opt := Ptr(&value)
			assert.Equal(t, 42, opt.AsValue())
		})

		t.Run("when value is nil", func(t *testing.T) {
			var value *int
			opt := Ptr(value)
			assert.True(t, opt.IsNil())
		})
	})

	t.Run("AndThen", func(t *testing.T) {
		t.Run("AndThen on Value Option", func(t *testing.T) {
			input := Value(5)
			mapper := func(x int) Option[int] {
				return Value(x * 2)
			}
			expected := Value(10)

			result := input.AndThen(mapper)

			assert.True(t, result.IsValue())
			assert.Equal(t, expected.AsValue(), result.AsValue())
		})

		t.Run("AndThen on Nil Option", func(t *testing.T) {
			input := Nil[int]()
			mapper := func(x int) Option[int] {
				return Value(x * 2)
			}

			result := input.AndThen(mapper)
			assert.True(t, result.IsNil())
		})
	})

	t.Run("IsValueAnd", func(t *testing.T) {
		t.Run("returns true for Value and a true predicate", func(t *testing.T) {
			input := Value(10)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsValueAnd(predicate)
			assert.True(t, result)
		})

		t.Run("returns false for Value and a false predicate", func(t *testing.T) {
			input := Value(3)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsValueAnd(predicate)
			assert.False(t, result)
		})

		t.Run("returns false for Nil", func(t *testing.T) {
			input := Nil[int]()
			predicate := func(x int) bool {
				return x > 5
			}

			result := input.IsValueAnd(predicate)
			assert.False(t, result)
		})
	})

	t.Run("IsNilOr", func(t *testing.T) {
		t.Run("returns true for Nil", func(t *testing.T) {
			input := Nil[int]()
			predicate := func(x int) bool {
				return x > 5
			}

			result := input.IsNilOr(predicate)
			assert.True(t, result)
		})

		t.Run("returns true for Value and a true predicate", func(t *testing.T) {
			input := Value(10)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsNilOr(predicate)
			assert.True(t, result)
		})

		t.Run("returns false for Value and a false predicate", func(t *testing.T) {
			input := Value(3)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsNilOr(predicate)
			assert.False(t, result)
		})
	})

	t.Run("Take", func(t *testing.T) {
		t.Run("Take on Value Option returns a Value and leaves a Nil", func(t *testing.T) {
			input := Value(10)
			expectedOld := Value(10)

			result := input.Take()
			assert.True(t, result.IsValue())
			assert.Equal(t, expectedOld.AsValue(), result.AsValue())
			assert.True(t, input.IsNil(), "The original option should be Nil after Take")
		})

		t.Run("Take on Nil Option returns a Nil and remains Nil", func(t *testing.T) {
			input := Nil[int]()

			result := input.Take()
			assert.True(t, result.IsNil())
			assert.True(t, input.IsNil(), "The original option should remain Nil after Take")
		})
	})

	t.Run("TakeIf", func(t *testing.T) {
		t.Run("TakeIf on a Value Option with a true predicate", func(t *testing.T) {
			input := Value(10)
			predicate := func(x int) bool { return x > 5 }

			result := input.TakeIf(predicate)
			assert.True(t, result.IsValue())
			assert.Equal(t, 10, result.AsValue())
			assert.True(t, input.IsNil(), "The original option should be Nil")
		})

		t.Run("TakeIf on a Value Option with a false predicate", func(t *testing.T) {
			input := Value(3)
			predicate := func(x int) bool { return x > 5 }

			result := input.TakeIf(predicate)
			assert.True(t, result.IsNil())
			assert.True(t, input.IsValue(), "The original option should remain Value")
		})

		t.Run("TakeIf on a Nil Option", func(t *testing.T) {
			input := Nil[int]()
			predicate := func(x int) bool {
				return x > 5
			}

			result := input.TakeIf(predicate)
			assert.True(t, result.IsNil())
			assert.True(t, input.IsNil(), "The original option should remain Nil")
		})
	})

	t.Run("OrPanic", func(t *testing.T) {
		t.Run("OrPanic on Value Option returns the value", func(t *testing.T) {
			input := Value(10)
			message := "This should not be printed"

			result := input.OrPanic(message)
			assert.Equal(t, 10, result)
		})

		t.Run("OrPanic on Nil Option panics with the message", func(t *testing.T) {
			input := Nil[int]()
			message := "Expected a value, but got nothing"

			assert.PanicsWithValue(t, message, func() {
				input.OrPanic(message)
			})
		})
	})

	t.Run("Insert", func(t *testing.T) {
		t.Run("Insert on a Nil Option", func(t *testing.T) {
			opt := Nil[int]()
			newValue := 42

			opt.Insert(newValue)

			assert.True(t, opt.IsValue(), "Option should become Value")
			assert.Equal(t, newValue, opt.AsValue(), "The inserted value should be correct")
		})

		t.Run("Insert on a Value Option", func(t *testing.T) {
			initialValue := 10
			opt := Value(initialValue)
			newValue := 20

			opt.Insert(newValue)

			assert.True(t, opt.IsValue(), "Option should remain Value")
			assert.Equal(t, newValue, opt.AsValue(), "The value should be updated")
		})
	})
}
