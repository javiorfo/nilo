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
	t.Run("Unwrap", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, 42, opt.Unwrap())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.Panics(t, func() {
				opt.Unwrap()
			})
		})
	})

	t.Run("UnwrapOr", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, 42, opt.UnwrapOr(24))
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.Equal(t, 24, opt.UnwrapOr(24))
		})
	})

	t.Run("UnwrapOrDefault", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(testStruct{"ok"})
			assert.Equal(t, "ok", opt.UnwrapOrDefault().Property)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[testStruct]()
			assert.Equal(t, "Default", opt.UnwrapOrDefault().Property)
		})
	})

	t.Run("OkOr", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			value, err := opt.OkOr(errors.New("error"))
			assert.Equal(t, 42, *value)
			assert.NoError(t, err)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			value, err := opt.OkOr(errors.New("error"))
			assert.Error(t, err)
			assert.Nil(t, value)
		})
	})

	t.Run("OkOrElse", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			value, err := opt.OkOrElse(func() error { return errors.New("error") })
			assert.Equal(t, 42, *value)
			assert.NoError(t, err)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			value, err := opt.OkOrElse(func() error { return errors.New("error") })
			assert.Error(t, err)
			assert.Nil(t, value)
		})
	})

	t.Run("OrElse", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, 42, opt.OrElse(func() Option[int] {
				return Some(24)
			}).Unwrap())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.Equal(t, 24, opt.OrElse(func() Option[int] {
				return Some(24)
			}).Unwrap())
		})
	})

	t.Run("Filter", func(t *testing.T) {
		t.Run("when value satisfies the filter", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, 42, opt.Filter(func(i int) bool {
				return i > 0
			}).Unwrap())
		})

		t.Run("when value does not satisfy the filter", func(t *testing.T) {
			opt := Some(42)
			assert.True(t, opt.Filter(func(i int) bool {
				return i < 0
			}).IsNone())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.True(t, opt.Filter(func(i int) bool {
				return i > 0
			}).IsNone())
		})
	})

	t.Run("IsNone", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.False(t, opt.IsNone())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.True(t, opt.IsNone())
		})
	})

	t.Run("IsSome", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.True(t, opt.IsSome())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.False(t, opt.IsSome())
		})
	})

	t.Run("Inspect", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			var result int
			opt.Inspect(func(i int) {
				result = i
			})
			assert.Equal(t, 42, result)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			var result int
			opt.Inspect(func(i int) {
				result = i
			})
			assert.Zero(t, result)
		})
	})

	t.Run("InspectOrElse", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			var result int
			opt.InspectOrElse(func(i int) {
				result = i
			}, func() {
				result = 24
			})
			assert.Equal(t, 42, result)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			var result int
			opt.InspectOrElse(func(i int) {
				result = i
			}, func() {
				result = 24
			})
			assert.Equal(t, 24, result)
		})
	})

	t.Run("UnwrapOrElse", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, 42, opt.UnwrapOrElse(func() int {
				return 24
			}))
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.Equal(t, 24, opt.UnwrapOrElse(func() int {
				return 24
			}))
		})
	})

	t.Run("None", func(t *testing.T) {
		opt := None[int]()
		assert.True(t, opt.IsNone())
	})

	t.Run("Some", func(t *testing.T) {
		opt := Some(42)
		assert.Equal(t, 42, opt.Unwrap())
	})

	t.Run("SomePtr", func(t *testing.T) {
		t.Run("when value is not nil", func(t *testing.T) {
			value := 42
			opt := SomePtr(&value)
			assert.Equal(t, 42, opt.Unwrap())
		})

		t.Run("when value is nil", func(t *testing.T) {
			var value *int
			opt := SomePtr(value)
			assert.True(t, opt.IsNone())
		})
	})

	t.Run("AndThen", func(t *testing.T) {
		t.Run("AndThen on Some Option", func(t *testing.T) {
			input := Some(5)
			mapper := func(x int) Option[int] {
				return Some(x * 2)
			}
			expected := Some(10)

			result := input.AndThen(mapper)

			assert.True(t, result.IsSome())
			assert.Equal(t, expected.Unwrap(), result.Unwrap())
		})

		t.Run("AndThen on None Option", func(t *testing.T) {
			input := None[int]()
			mapper := func(x int) Option[int] {
				return Some(x * 2)
			}

			result := input.AndThen(mapper)
			assert.True(t, result.IsNone())
		})
	})

	t.Run("And", func(t *testing.T) {
		t.Run("And with two Some Options", func(t *testing.T) {
			a := Some(10)
			b := Some(20)
			expected := Some(20)

			result := a.And(b)
			assert.True(t, result.IsSome())
			assert.Equal(t, expected.Unwrap(), result.Unwrap())
		})

		t.Run("And with a Some and a None", func(t *testing.T) {
			a := Some(10)
			b := None[int]()

			result := a.And(b)
			assert.True(t, result.IsNone())
		})

		t.Run("And with a None and a Some", func(t *testing.T) {
			a := None[int]()
			b := Some(20)

			result := a.And(b)
			assert.True(t, result.IsNone())
		})

		t.Run("And with two None Options", func(t *testing.T) {
			a := None[int]()
			b := None[int]()

			result := a.And(b)
			assert.True(t, result.IsNone())
		})
	})

	t.Run("Or", func(t *testing.T) {
		t.Run("Or with two Some Options", func(t *testing.T) {
			a := Some(10)
			b := Some(20)

			result := a.Or(b)

			assert.True(t, result.IsSome())
			assert.Equal(t, 10, result.Unwrap())
		})

		t.Run("Or with a Some and a None", func(t *testing.T) {
			a := Some(10)
			b := None[int]()

			result := a.Or(b)

			assert.True(t, result.IsSome())
			assert.Equal(t, 10, result.Unwrap())
		})

		t.Run("Or with a None and a Some", func(t *testing.T) {
			a := None[int]()
			b := Some(20)

			result := a.Or(b)

			assert.True(t, result.IsSome())
			assert.Equal(t, 20, result.Unwrap())
		})

		t.Run("Or with two None Options", func(t *testing.T) {
			a := None[int]()
			b := None[int]()

			result := a.Or(b)

			assert.True(t, result.IsNone())
		})
	})

	t.Run("Xor", func(t *testing.T) {
		t.Run("Xor with two Some Options returns None", func(t *testing.T) {
			a := Some(10)
			b := Some(20)

			result := a.Xor(b)

			assert.True(t, result.IsNone())
		})

		t.Run("Xor with Some and None returns Some", func(t *testing.T) {
			a := Some(10)
			b := None[int]()

			result := a.Xor(b)

			assert.True(t, result.IsSome())
			assert.Equal(t, 10, result.Unwrap())
		})

		t.Run("Xor with None and Some returns Some", func(t *testing.T) {
			a := None[int]()
			b := Some(20)

			result := a.Xor(b)

			assert.True(t, result.IsSome())
			assert.Equal(t, 20, result.Unwrap())
		})

		t.Run("Xor with two None Options returns None", func(t *testing.T) {
			a := None[int]()
			b := None[int]()

			result := a.Xor(b)

			assert.True(t, result.IsNone())
		})
	})

	t.Run("IsSomeAnd", func(t *testing.T) {
		t.Run("IsSomeAnd returns true for Some and a true predicate", func(t *testing.T) {
			input := Some(10)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsSomeAnd(predicate)
			assert.True(t, result)
		})

		t.Run("IsSomeAnd returns false for Some and a false predicate", func(t *testing.T) {
			input := Some(3)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsSomeAnd(predicate)
			assert.False(t, result)
		})

		t.Run("IsSomeAnd returns false for None", func(t *testing.T) {
			input := None[int]()
			predicate := func(x int) bool {
				return x > 5
			}

			result := input.IsSomeAnd(predicate)
			assert.False(t, result)
		})
	})

	t.Run("IsNoneOr", func(t *testing.T) {
		t.Run("IsNoneOr returns true for None", func(t *testing.T) {
			input := None[int]()
			predicate := func(x int) bool {
				return x > 5
			}

			result := input.IsNoneOr(predicate)
			assert.True(t, result)
		})

		t.Run("IsNoneOr returns true for Some and a true predicate", func(t *testing.T) {
			input := Some(10)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsNoneOr(predicate)
			assert.True(t, result)
		})

		t.Run("IsNoneOr returns false for Some and a false predicate", func(t *testing.T) {
			input := Some(3)
			predicate := func(x int) bool { return x > 5 }

			result := input.IsNoneOr(predicate)
			assert.False(t, result)
		})
	})

	t.Run("Take", func(t *testing.T) {
		t.Run("Take on Some Option returns a Some and leaves a None", func(t *testing.T) {
			input := Some(10)
			expectedOld := Some(10)

			result := input.Take()
			assert.True(t, result.IsSome())
			assert.Equal(t, expectedOld.Unwrap(), result.Unwrap())
			assert.True(t, input.IsNone(), "The original option should be None after Take")
		})

		t.Run("Take on None Option returns a None and remains None", func(t *testing.T) {
			input := None[int]()

			result := input.Take()
			assert.True(t, result.IsNone())
			assert.True(t, input.IsNone(), "The original option should remain None after Take")
		})
	})

	t.Run("TakeIf", func(t *testing.T) {
		t.Run("TakeIf on a Some Option with a true predicate", func(t *testing.T) {
			input := Some(10)
			predicate := func(x int) bool { return x > 5 }

			result := input.TakeIf(predicate)
			assert.True(t, result.IsSome())
			assert.Equal(t, 10, result.Unwrap())
			assert.True(t, input.IsNone(), "The original option should be None")
		})

		t.Run("TakeIf on a Some Option with a false predicate", func(t *testing.T) {
			input := Some(3)
			predicate := func(x int) bool { return x > 5 }

			result := input.TakeIf(predicate)
			assert.True(t, result.IsNone())
			assert.True(t, input.IsSome(), "The original option should remain Some")
		})

		t.Run("TakeIf on a None Option", func(t *testing.T) {
			input := None[int]()
			predicate := func(x int) bool {
				return x > 5
			}

			result := input.TakeIf(predicate)
			assert.True(t, result.IsNone())
			assert.True(t, input.IsNone(), "The original option should remain None")
		})
	})

	t.Run("Replace", func(t *testing.T) {
		t.Run("Replace on Some Option", func(t *testing.T) {
			input := Some(10)
			newValue := 20

			opt := input.Replace(newValue)
			assert.True(t, opt.IsSome(), "The returned option should be Some")
			assert.Equal(t, 20, input.Unwrap(), "The original option should have the new value")
		})

		t.Run("Replace on None Option", func(t *testing.T) {
			input := None[int]()
			newValue := 20

			returnedOption := input.Replace(newValue)
			assert.True(t, returnedOption.IsSome(), "The returned option should be Some")
			assert.Equal(t, 20, input.Unwrap(), "The original option should have the new value")
		})
	})

	t.Run("Expect", func(t *testing.T) {
		t.Run("Expect on Some Option returns the value", func(t *testing.T) {
			input := Some(10)
			message := "This should not be printed"

			result := input.Expect(message)
			assert.Equal(t, 10, result)
		})

		t.Run("Expect on None Option panics with the message", func(t *testing.T) {
			input := None[int]()
			message := "Expected a value, but got nothing"

			assert.PanicsWithValue(t, message, func() {
				input.Expect(message)
			})
		})
	})

	t.Run("Insert", func(t *testing.T) {
		t.Run("Insert on a None Option", func(t *testing.T) {
			opt := None[int]()
			newValue := 42

			opt.Insert(newValue)

			assert.True(t, opt.IsSome(), "Option should become Some")
			assert.Equal(t, newValue, opt.Unwrap(), "The inserted value should be correct")
		})

		t.Run("Insert on a Some Option", func(t *testing.T) {
			initialValue := 10
			opt := Some(initialValue)
			newValue := 20

			opt.Insert(newValue)

			assert.True(t, opt.IsSome(), "Option should remain Some")
			assert.Equal(t, newValue, opt.Unwrap(), "The value should be updated")
		})
	})

	t.Run("GetOrInsert", func(t *testing.T) {
		t.Run("GetOrInsert on a Some Option returns the existing value", func(t *testing.T) {
			initialValue := 10
			opt := Some(initialValue)

			returnedValue := opt.GetOrInsert(999)
			assert.True(t, opt.IsSome(), "Option should remain Some")
			assert.Equal(t, initialValue, opt.Unwrap(), "Value should not be changed")
			assert.Equal(t, initialValue, returnedValue, "Returned value should be the original value")
		})

		t.Run("GetOrInsert on a None Option inserts and returns the new value", func(t *testing.T) {
			opt := None[int]()
			expectedValue := 42

			returnedValue := opt.GetOrInsert(expectedValue)
			assert.True(t, opt.IsSome(), "Option should become Some")
			assert.Equal(t, expectedValue, opt.Unwrap(), "Value should be correctly inserted")
			assert.Equal(t, expectedValue, returnedValue, "Returned value should be the newly inserted value")
		})
	})

	t.Run("GetOrInsertWith", func(t *testing.T) {
		t.Run("GetOrInsertWith on a Some Option returns the existing value", func(t *testing.T) {
			initialValue := 10
			opt := Some(initialValue)

			returnedValue := opt.GetOrInsertWith(func() int {
				return 999
			})

			assert.True(t, opt.IsSome(), "Option should remain Some")
			assert.Equal(t, initialValue, opt.Unwrap(), "Value should not be changed")
			assert.Equal(t, initialValue, returnedValue, "Returned value should be the original value")
		})

		t.Run("GetOrInsertWith on a None Option inserts and returns a new value", func(t *testing.T) {
			opt := None[int]()
			expectedValue := 42

			returnedValue := opt.GetOrInsertWith(func() int {
				return expectedValue
			})

			assert.True(t, opt.IsSome(), "Option should become Some")
			assert.Equal(t, expectedValue, opt.Unwrap(), "Value should be correctly inserted")
			assert.Equal(t, expectedValue, returnedValue, "Returned value should be the newly inserted value")
		})
	})

	t.Run("GetOrInsertDefault", func(t *testing.T) {
		t.Run("GetOrInsertDefault on a Some Option returns existing value", func(t *testing.T) {
			initialValue := 10
			opt := Some(initialValue)

			returnedValue := opt.GetOrInsertDefault()

			assert.True(t, opt.IsSome(), "Option should remain Some")
			assert.Equal(t, initialValue, opt.Unwrap(), "Value should not be changed")
			assert.Equal(t, initialValue, returnedValue, "Returned value should be the original value")
		})

		t.Run("GetOrInsertDefault on a None Option returns zero value for built-in type", func(t *testing.T) {
			opt := None[int]()

			returnedValue := opt.GetOrInsertDefault()

			assert.True(t, opt.IsSome(), "Option should become Some")
			assert.Equal(t, 0, opt.Unwrap(), "Value should be the zero value")
			assert.Equal(t, 0, returnedValue, "Returned value should be the zero value")
		})

		t.Run("GetOrInsertDefault on a None Option returns custom default value", func(t *testing.T) {
			opt := None[testStruct]()

			returnedValue := opt.GetOrInsertDefault()

			assert.True(t, opt.IsSome(), "Option should become Some")
			assert.Equal(t, "Default", opt.Unwrap().Property, "Value should be the custom default value")
			assert.Equal(t, "Default", returnedValue.Property, "Returned value should be the custom default value")
		})
	})
}
