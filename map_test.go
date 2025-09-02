package nilo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Map", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, 84, opt.Map(func(i int) int {
				return i * 2
			}).Unwrap())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.True(t, opt.Map(func(i int) int {
				return i * 2
			}).IsNone())
		})
	})

	t.Run("MapToAny", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, true, opt.MapToAny(func(i int) any {
				return i > 0
			}).Unwrap())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.True(t, opt.MapToAny(func(i int) any {
				return i > 0
			}).IsNone())
		})
	})

	t.Run("MapToString", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Some(42)
			assert.Equal(t, "Value 42", opt.MapToString(func(i int) string {
				return fmt.Sprintf("Value %d", i)
			}).Unwrap())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := None[int]()
			assert.True(t, opt.MapToString(func(i int) string {
				return fmt.Sprintf("Value %d", i)
			}).IsNone())
		})
	})

	t.Run("MapOr", func(t *testing.T) {
		t.Run("MapOr on a Some Option returns the mapped value", func(t *testing.T) {
			input := Some(5)
			defaultValue := 100
			mapper := func(x int) int {
				return x * 2
			}
			expected := 10

			result := input.MapOr(defaultValue, mapper)
			assert.Equal(t, expected, result)
		})

		t.Run("MapOr on a None Option returns the default value", func(t *testing.T) {
			input := None[int]()
			defaultValue := 100
			mapper := func(x int) int {
				return x * 2
			}
			expected := 100

			result := input.MapOr(defaultValue, mapper)
			assert.Equal(t, expected, result)
		})
	})

	t.Run("MapOrAny", func(t *testing.T) {
		t.Run("MapOrAny on a Some Option returns the mapped value", func(t *testing.T) {
			input := Some(5)
			defaultValue := 100
			mapper := func(x int) any {
				return x * 2
			}
			expected := 10

			result := input.MapOrAny(defaultValue, mapper)
			assert.Equal(t, expected, result)
		})

		t.Run("MapOrAny on a None Option returns the default value", func(t *testing.T) {
			input := None[int]()
			defaultValue := 100
			mapper := func(x int) any {
				return x * 2
			}
			expected := 100

			result := input.MapOrAny(defaultValue, mapper)
			assert.Equal(t, expected, result)
		})
	})

	t.Run("MapOrElse", func(t *testing.T) {
		t.Run("MapOrElse on a Some Option returns the mapped value", func(t *testing.T) {
			input := Some(5)
			supplier := func() int {
				return 100 // This should not be called
			}
			mapper := func(x int) int {
				return x * 2
			}
			expected := 10

			result := input.MapOrElse(supplier, mapper)
			assert.Equal(t, expected, result)
		})

		t.Run("MapOrElse on a None Option returns the value from the supplier", func(t *testing.T) {
			input := None[int]()
			supplier := func() int {
				return 100 // This should be called
			}
			mapper := func(x int) int {
				return x * 2
			}
			expected := 100

			result := input.MapOrElse(supplier, mapper)
			assert.Equal(t, expected, result)
		})
	})

	t.Run("MapOrDefault", func(t *testing.T) {
		t.Run("MapOrDefault on a Some Option returns the mapped value", func(t *testing.T) {
			input := Some(5)
			mapper := func(x int) int {
				return x * 2
			}

			result := input.MapOrDefault(mapper)
			assert.Equal(t, 10, result)
		})

		t.Run("MapOrDefault on a None Option with a built-in type returns the zero value", func(t *testing.T) {
			input := None[int]()
			mapper := func(x int) int {
				return x * 2
			}

			result := input.MapOrDefault(mapper)

			assert.Equal(t, 0, result)
		})

		t.Run("MapOrDefault on a None Option with a custom type returns the custom default", func(t *testing.T) {
			input := None[testStruct]()
			mapper := func(m testStruct) testStruct {
				return testStruct{Property: "Def"}
			}

			result := input.MapOrDefault(mapper)
			assert.Equal(t, "Default", result.Property)
		})
	})
}
