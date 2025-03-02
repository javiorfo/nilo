package nilo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptional(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.Equal(t, 42, opt.Get())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.Panics(t, func() {
				opt.Get()
			})
		})
	})

	t.Run("OrElse", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.Equal(t, 42, opt.OrElse(24))
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.Equal(t, 24, opt.OrElse(24))
		})
	})

	t.Run("OrError", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			value, err := opt.OrError(errors.New("error"))
			assert.Equal(t, 42, *value)
			assert.NoError(t, err)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			value, err := opt.OrError(errors.New("error"))
			assert.Error(t, err)
			assert.Nil(t, value)
		})
	})

	t.Run("Or", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.Equal(t, 42, opt.Or(func() Optional[int] {
				return From(24)
			}).Get())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.Equal(t, 24, opt.Or(func() Optional[int] {
				return From(24)
			}).Get())
		})
	})

	t.Run("Filter", func(t *testing.T) {
		t.Run("when value satisfies the filter", func(t *testing.T) {
			opt := From(42)
			assert.Equal(t, 42, opt.Filter(func(i int) bool {
				return i > 0
			}).Get())
		})

		t.Run("when value does not satisfy the filter", func(t *testing.T) {
			opt := From(42)
			assert.True(t, opt.Filter(func(i int) bool {
				return i < 0
			}).IsEmpty())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.True(t, opt.Filter(func(i int) bool {
				return i > 0
			}).IsEmpty())
		})
	})

	t.Run("MapToAny", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.Equal(t, 84, opt.MapToAny(func(i int) any {
				return i * 2
			}).Get())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.True(t, opt.MapToAny(func(i int) any {
				return i * 2
			}).IsEmpty())
		})
	})

	t.Run("MapToString", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.Equal(t, "Value 42", opt.MapToString(func(i int) string {
				return fmt.Sprintf("Value %d", i)
			}).Get())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.True(t, opt.MapToString(func(i int) string {
				return fmt.Sprintf("Value %d", i)
			}).IsEmpty())
		})
	})

	t.Run("IsEmpty", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.False(t, opt.IsEmpty())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.True(t, opt.IsEmpty())
		})
	})

	t.Run("IsPresent", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.True(t, opt.IsPresent())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.False(t, opt.IsPresent())
		})
	})

	t.Run("IfPresent", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			var result int
			opt.IfPresent(func(i int) {
				result = i
			})
			assert.Equal(t, 42, result)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			var result int
			opt.IfPresent(func(i int) {
				result = i
			})
			assert.Zero(t, result)
		})
	})

	t.Run("IfPresentOrElse", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			var result int
			opt.IfPresentOrElse(func(i int) {
				result = i
			}, func() {
				result = 24
			})
			assert.Equal(t, 42, result)
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			var result int
			opt.IfPresentOrElse(func(i int) {
				result = i
			}, func() {
				result = 24
			})
			assert.Equal(t, 24, result)
		})
	})

	t.Run("OrElseGet", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			assert.Equal(t, 42, opt.OrElseGet(func() int {
				return 24
			}))
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			assert.Equal(t, 24, opt.OrElseGet(func() int {
				return 24
			}))
		})
	})

	t.Run("Empty", func(t *testing.T) {
		opt := Empty[int]()
		assert.True(t, opt.IsEmpty())
	})

	t.Run("From", func(t *testing.T) {
		opt := From(42)
		assert.Equal(t, 42, opt.Get())
	})

	t.Run("FromPtr", func(t *testing.T) {
		t.Run("when value is not nil", func(t *testing.T) {
			value := 42
			opt := FromPtr(&value)
			assert.Equal(t, 42, opt.Get())
		})

		t.Run("when value is nil", func(t *testing.T) {
			var value *int
			opt := FromPtr(value)
			assert.True(t, opt.IsEmpty())
		})
	})

	t.Run("Map", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := From(42)
			mapped := Map(opt, func(i int) string {
				return fmt.Sprintf("value: %d", i)
			})
			assert.Equal(t, "value: 42", mapped.Get())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Empty[int]()
			mapped := Map(opt, func(i int) string {
				return fmt.Sprintf("value: %d", i)
			})
			assert.True(t, mapped.IsEmpty())
		})
	})
}
