package nilo

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// MarshalJSON implements the `json.Marshaler` interface for `Option`.
//
// If the `Option` is `None`, it marshals to the JSON value `null`.
// If the `Option` is `Some`, it marshals the wrapped value to its JSON representation.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		return []byte("null"), nil
	}
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the `json.Unmarshaler` interface for `Option`.
//
// If the JSON data is `null`, it unmarshals into a `None` `Option`.
// Otherwise, it unmarshals the data into the `Option`'s value, creating a
// `Some` `Option` with the unmarshaled content.
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
// It returns a string representation of the `Option`. For `Some` `Option`s,
// the format is "Some(value)". For a `None` `Option`, the format is "None".
func (o Option[T]) String() string {
	if o.IsSome() {
		return fmt.Sprintf("Some(%v)", o.Unwrap())
	}
	return "None"
}
