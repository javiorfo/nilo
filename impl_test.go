package nilo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpl(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		t.Run("MarshalJSON on a Some Option", func(t *testing.T) {
			input := Some("hello")
			expectedJSON := `"hello"`

			result, err := json.Marshal(input)

			assert.NoError(t, err)
			assert.Equal(t, expectedJSON, string(result))
		})

		t.Run("MarshalJSON on a None Option", func(t *testing.T) {
			input := None[string]()
			expectedJSON := `null`

			result, err := json.Marshal(input)

			assert.NoError(t, err)
			assert.Equal(t, expectedJSON, string(result))
		})
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Run("UnmarshalJSON from null", func(t *testing.T) {
			var opt Option[string]
			jsonData := []byte("null")

			err := opt.UnmarshalJSON(jsonData)

			assert.NoError(t, err)
			assert.True(t, opt.IsNone(), "Option should be None after unmarshaling from null")
		})

		t.Run("UnmarshalJSON from a valid value", func(t *testing.T) {
			var opt Option[string]
			jsonData := []byte(`"hello"`)
			expectedValue := "hello"

			err := opt.UnmarshalJSON(jsonData)

			assert.NoError(t, err)
			assert.True(t, opt.IsSome(), "Option should be Some after unmarshaling from a value")
			assert.Equal(t, expectedValue, opt.Unwrap(), "The unmarshaled value should be correct")
		})

		t.Run("UnmarshalJSON with invalid data returns an error", func(t *testing.T) {
			var opt Option[int]
			jsonData := []byte(`"not an int"`)

			err := opt.UnmarshalJSON(jsonData)

			assert.Error(t, err, "Unmarshal should fail for invalid data")
			assert.True(t, opt.IsNone(), "Option should remain None on error")
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("String representation of a Some Option", func(t *testing.T) {
			input := Some(10)
			expected := "Some(10)"

			result := input.String()
			assert.Equal(t, expected, result)
		})

		t.Run("String representation of a None Option", func(t *testing.T) {
			input := None[int]()
			expected := "None"

			result := input.String()
			assert.Equal(t, expected, result)
		})
	})
}
