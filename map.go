package nilo

// Map applies a function to the contained value of an Option if it is `Some`
// and returns a new `Option` containing the mapped value.
//
// If the original `Option` is `None`, this method returns `None`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns a value of the same type `T`.
//
// Returns:
//   - A new `Option[T]` containing the result of the mapping, or `None` if the
//     original `Option` was `None`.
func (o Option[T]) Map(mapper func(T) T) Option[T] {
	if o.IsSome() {
		return Some(mapper(o.Unwrap()))
	}
	return None[T]()
}

// MapToString maps the contained value of an `Option[T]` to a string if it is `Some`,
// and returns a new `Option[string]` with the result.
//
// If the original `Option` is `None`, this method returns `None[string]`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns a `string`.
//
// Returns:
//   - A new `Option[string]` containing the mapped value, or `None[string]`
//     if the original `Option` was `None`.
func (o Option[T]) MapToString(mapper func(T) string) Option[string] {
	if o.IsSome() {
		return Some(mapper(o.Unwrap()))
	}
	return None[string]()
}

// MapToInt maps the contained value of an `Option[T]` to an integer if it is `Some`,
// and returns a new `Option[int]` with the result.
//
// If the original `Option` is `None`, this method returns `None[int]`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns an `int`.
//
// Returns:
//   - A new `Option[int]` containing the mapped value, or `None[int]`
//     if the original `Option` was `None`.
func (o Option[T]) MapToInt(mapper func(T) int) Option[int] {
	if o.IsSome() {
		return Some(mapper(o.Unwrap()))
	}
	return None[int]()
}

// MapToBool maps the contained value of an `Option[T]` to a boolean if it is `Some`,
// and returns a new `Option[bool]` with the result.
//
// If the original `Option` is `None`, this method returns `None[bool]`.
//
// Parameters:
//   - mapper: The function to apply to the `Option`'s value. It takes a value
//     of type `T` and returns a `bool`.
//
// Returns:
//   - A new `Option[bool]` containing the mapped value, or `None[bool]`
//     if the original `Option` was `None`.
func (o Option[T]) MapToBool(mapper func(T) bool) Option[bool] {
	if o.IsSome() {
		return Some(mapper(o.Unwrap()))
	}
	return None[bool]()
}

// MapOr maps the `Option`'s value if it is `Some` and returns the result,
// otherwise returns a default value.
//
// Parameters:
//   - def: The default value to return if the `Option` is `None`.
//   - mapper: A function to apply to the `Option`'s value if it is `Some`.
func (o Option[T]) MapOr(def T, mapper func(T) T) T {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return def
}

// MapOrString maps the `Option`'s value to a string if it is `Some` and
// returns the result, otherwise returns a default string.
//
// Parameters:
//   - def: The default string to return if the `Option` is `None`.
//   - mapper: A function to apply to the `Option`'s value to produce a string.
func (o Option[T]) MapOrString(def string, mapper func(T) string) string {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return def
}

// MapOrInt maps the `Option`'s value to an integer if it is `Some` and
// returns the result, otherwise returns a default integer.
//
// Parameters:
//   - def: The default integer to return if the `Option` is `None`.
//   - mapper: A function to apply to the `Option`'s value to produce an integer.
func (o Option[T]) MapOrInt(def int, mapper func(T) int) int {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return def
}

// MapOrBool maps the `Option`'s value to a boolean if it is `Some` and
// returns the result, otherwise returns a default boolean.
//
// Parameters:
//   - def: The default boolean to return if the `Option` is `None`.
//   - mapper: A function to apply to the `Option`'s value to produce a boolean.
func (o Option[T]) MapOrBool(def bool, mapper func(T) bool) bool {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return def
}

// MapOrElse maps the `Option`'s value if it is `Some` and returns the result,
// otherwise calls a supplier function to get the default value.
//
// Parameters:
//
//	  the function signature).
//	- supplier: A function that provides the default value if the `Option` is `None`.
//	- mapper: A function to apply to the `Option`'s value if it is `Some`.
func (o Option[T]) MapOrElse(supplier func() T, mapper func(T) T) T {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return supplier()
}

// MapOrElseString maps the `Option`'s value to a string if it is `Some`,
// otherwise calls a supplier function to get the default string.
//
// Parameters:
//
//	  the function signature).
//	- supplier: A function that provides the default string if the `Option` is `None`.
//	- mapper: A function to apply to the `Option`'s value to produce a string.
func (o Option[T]) MapOrElseString(supplier func() string, mapper func(T) string) string {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return supplier()
}

// MapOrElseInt maps the `Option`'s value to an integer if it is `Some`,
// otherwise calls a supplier function to get the default integer.
//
// Parameters:
//
//	  the function signature).
//	- supplier: A function that provides the default integer if the `Option` is `None`.
//	- mapper: A function to apply to the `Option`'s value to produce an integer.
func (o Option[T]) MapOrElseInt(supplier func() int, mapper func(T) int) int {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return supplier()
}

// MapOrElseBool maps the `Option`'s value to a boolean if it is `Some`,
// otherwise calls a supplier function to get the default boolean.
//
// Parameters:
//
//	  the function signature).
//	- supplier: A function that provides the default boolean if the `Option` is `None`.
//	- mapper: A function to apply to the `Option`'s value to produce a boolean.
func (o Option[T]) MapOrElseBool(supplier func() bool, mapper func(T) bool) bool {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}
	return supplier()
}

// MapOrDefault maps the `Option`'s value if it is `Some`, otherwise returns
// the zero value of the type.
//
// Parameters:
//   - mapper: A function to apply to the `Option`'s value if it is `Some`.
func (o Option[T]) MapOrDefault(mapper func(T) T) T {
	if o.IsSome() {
		return mapper(o.Unwrap())
	}

	return defaultImplOrNew[T]()
}
