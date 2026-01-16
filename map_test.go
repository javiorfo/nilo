package nilo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Map", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			assert.Equal(t, 84, opt.Map(func(i int) int {
				return i * 2
			}).AsValue())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.True(t, opt.Map(func(i int) int {
				return i * 2
			}).IsNil())
		})
	})

	t.Run("MapToString", func(t *testing.T) {
		t.Run("when value is present", func(t *testing.T) {
			opt := Value(42)
			assert.Equal(t, "Value 42", opt.MapToString(func(i int) string {
				return fmt.Sprintf("Value %d", i)
			}).AsValue())
		})

		t.Run("when value is not present", func(t *testing.T) {
			opt := Nil[int]()
			assert.True(t, opt.MapToString(func(i int) string {
				return fmt.Sprintf("Value %d", i)
			}).IsNil())
		})
	})

	t.Run("MapOrDefault", func(t *testing.T) {
		t.Run("MapOrDefault on a Value Option returns the mapped value", func(t *testing.T) {
			input := Value(5)
			mapper := func(x int) int {
				return x * 2
			}

			result := input.MapOrDefault(mapper)
			assert.Equal(t, 10, result)
		})

		t.Run("MapOrDefault on a Nil Option with a built-in type returns the zero value", func(t *testing.T) {
			input := Nil[int]()
			mapper := func(x int) int {
				return x * 2
			}

			result := input.MapOrDefault(mapper)

			assert.Equal(t, 0, result)
		})

		t.Run("MapOrDefault on a Nil Option with a custom type returns the custom default", func(t *testing.T) {
			input := Nil[testStruct]()
			mapper := func(m testStruct) testStruct {
				return testStruct{Property: "Def"}
			}

			result := input.MapOrDefault(mapper)
			assert.Equal(t, "Default", result.Property)
		})
	})
}
