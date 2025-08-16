package nilo

// FromResult creates an `Option` from a Go function's return values.
//
// It returns a `Some` `Option` containing `value` if `err` is `nil`.
// If `err` is not `nil`, it returns a `None` `Option`.
// This is a convenient way to bridge error-handling patterns in Go with the
// `Option` type.
//
// Parameters:
//   - value: The value to wrap in a `Some` `Option` if there is no error.
//   - err: The error returned from a function.
func FromResult[T any](value T, err error) Option[T] {
	if err != nil {
		return None[T]()
	}
	return Some(value)
}

// OkAndResult applies a function that returns a value and an error to the
// `Option`'s contained value.
//
// If the `Option` is `Some` and the applied function returns a `nil` error,
// this method returns a `Some` `Option` with the function's returned value.
// In all other cases (if the `Option` is `None` or the function returns a
// non-nil error), it returns a `None` `Option`.
// This is useful for chaining operations that might fail.
//
// Parameters:
//   - apply: A function that takes the `Option`'s value and returns a new value
//     and an error.
func (o Option[T]) OkAndResult(apply func(T) (T, error)) Option[T] {
	if o.IsSome() {
		if r, err := apply(o.Unwrap()); err == nil {
			return Some(r)
		}
	}
	return None[T]()
}
