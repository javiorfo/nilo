package nilo

// FromResult creates an `Option` from a Go function's return values.
//
// It returns a `Value` `Option` containing `value` if `err` is `nil`.
// If `err` is not `nil`, it returns a `Nil` `Option`.
// This is a convenient way to bridge error-handling patterns in Go with the
// `Option` type.
//
// Parameters:
//   - value: The value to wrap in a `Value` `Option` if there is no error.
//   - err: The error returned from a function.
func FromResult[T any](value T, err error) Option[T] {
	if err != nil {
		return Nil[T]()
	}
	return Value(value)
}

// AndResult applies a function that returns a value and an error to the
// `Option`'s contained value.
//
// If the `Option` is `Value` and the applied function returns a `nil` error,
// this method returns a `Value` `Option` with the function's returned value.
// In all other cases (if the `Option` is `Nil` or the function returns a
// non-nil error), it returns a `Nil` `Option`.
// This is useful for chaining operations that might fail.
//
// Parameters:
//   - apply: A function that takes the `Option`'s value and returns a new value
//     and an error.
func (o Option[T]) AndResult(apply func(T) (T, error)) Option[T] {
	if o.IsValue() {
		if r, err := apply(o.AsValue()); err == nil {
			return Value(r)
		}
	}
	return Nil[T]()
}

// AndPtrResult applies a function that returns a pointer value and an error to the
// `Option`'s contained value.
//
// If the `Option` is `Value` and the applied function returns a `nil` error,
// this method returns a `Value` `Option` with the function's returned value.
// In all other cases (if the `Option` is `Nil` or the function returns a
// non-nil error or a nil value), it returns a `Nil` `Option`.
// This is useful for chaining operations that might fail.
//
// Parameters:
//   - apply: A function that takes the `Option`'s value and returns a new
//     pointer value and an error.
func (o Option[T]) AndPtrResult(apply func(T) (*T, error)) Option[T] {
	if o.IsValue() {
		if r, err := apply(o.AsValue()); err == nil && r != nil {
			return Value(*r)
		}
	}
	return Nil[T]()
}
