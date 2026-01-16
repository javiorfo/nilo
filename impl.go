package nilo

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// MarshalJSON implements the `json.Marshaler` interface for `Option`.
//
// If the `Option` is `Nil`, it marshals to the JSON value `null`.
// If the `Option` is `Value`, it marshals the wrapped value to its JSON representation.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNil() {
		return []byte("null"), nil
	}
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the `json.Unmarshaler` interface for `Option`.
//
// If the JSON data is `null`, it unmarshals into a `Nil` `Option`.
// Otherwise, it unmarshals the data into the `Option`'s value, creating a
// `Value` `Option` with the unmarshaled content.
func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		o.value = nil
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	o.value = &v
	return nil
}

// String implements the `fmt.Stringer` interface for `Option`.
//
// It returns a string representation of the `Option`. For `Value` `Option`s,
// the format is "Value". For a `Nil` `Option`, the format is "Nil".
func (o Option[T]) String() string {
	if o.IsValue() {
		return fmt.Sprintf("Value(%v)", o.AsValue())
	}
	return "Nil"
}
