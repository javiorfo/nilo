package nilo

// Map applies a function to the contained value of an Option if it is `Value`
// and returns a new `Option` containing the mapped value.
//
// If the original `Option` is `Nil`, this method returns `Nil`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns a value of the same type `T`.
//
// Returns:
//   - A new `Option[T]` containing the result of the mapping, or `Nil` if the
//     original `Option` was `Nil`.
func (o Option[T]) Map(mapper func(T) T) Option[T] {
	if o.IsValue() {
		return Value(mapper(o.AsValue()))
	}
	return Nil[T]()
}

// MapToString maps the contained value of an `Option[T]` to a string if it is `Value`,
// and returns a new `Option[string]` with the result.
//
// If the original `Option` is `Nil`, this method returns `Nil string`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns a `string`.
//
// Returns:
//   - A new `Option[string]` containing the mapped value, or `Nil string`
//     if the original `Option` was `Nil`.
func (o Option[T]) MapToString(mapper func(T) string) Option[string] {
	if o.IsValue() {
		return Value(mapper(o.AsValue()))
	}
	return Nil[string]()
}

// MapToInt maps the contained value of an `Option[T]` to an integer if it is `Value`,
// and returns a new `Option[int]` with the result.
//
// If the original `Option` is `Nil`, this method returns `Nil int`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns an `int`.
//
// Returns:
//   - A new `Option[int]` containing the mapped value, or `Nil int`
//     if the original `Option` was `Nil`.
func (o Option[T]) MapToInt(mapper func(T) int) Option[int] {
	if o.IsValue() {
		return Value(mapper(o.AsValue()))
	}
	return Nil[int]()
}

// MapToBool maps the contained value of an `Option[T]` to a boolean if it is `Value`,
// and returns a new `Option[bool]` with the result.
//
// If the original `Option` is `Nil`, this method returns `Nil bool`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns a `bool`.
//
// Returns:
//   - A new `Option[bool]` containing the mapped value, or `Nil bool`
//     if the original `Option` was `Nil`.
func (o Option[T]) MapToBool(mapper func(T) bool) Option[bool] {
	if o.IsValue() {
		return Value(mapper(o.AsValue()))
	}
	return Nil[bool]()
}

// MapOrDefault maps the `Option`'s value if it is `Value`, otherwise returns
// the zero value of the type or the implemented in Default interface.
//
// Parameters:
//   - mapper: A function to apply to the `Option`'s value if it is `Value`.
func (o Option[T]) MapOrDefault(mapper func(T) T) T {
	if o.IsValue() {
		return mapper(o.AsValue())
	}

	return defaultImplOrNew[T]()
}
