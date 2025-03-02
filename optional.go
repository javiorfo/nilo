// Package nilo provides a generic Optional type that can be used to represent a value that may or may not be present.
package nilo

// Optional is a generic type that encapsulates a value that may or may not be present.
// It provides methods to work with the value safely.
type Optional[T any] struct {
	value *T
}

// Get retrieves the value contained in the Optional.
// It panics if the value is not present (i.e., if IsEmpty() returns true).
func (o Optional[T]) Get() T {
	return *o.value
}

// OrElse returns the value contained in the Optional if present; otherwise, it returns the provided alternative value.
func (o Optional[T]) OrElse(other T) T {
	if o.IsPresent() {
		return o.Get()
	}
	return other
}

// OrError returns the value contained in the Optional if present; otherwise, it returns the provided error.
// This returns pointer type or an error
func (o Optional[T]) OrError(err error) (*T, error) {
	if o.IsPresent() {
		return o.value, nil
	}
	return nil, err
}

// Or returns the value contained in the Optional if present; otherwise, it invokes the provided supplier function
// to obtain a new Optional.
func (o Optional[T]) Or(supplier func() Optional[T]) Optional[T] {
	if o.IsEmpty() {
		return supplier()
	}
	return Of(o.Get())
}

// Filter returns an Optional containing the value if it is present and satisfies the provided filter function;
// otherwise, it returns an empty Optional.
func (o Optional[T]) Filter(filter func(T) bool) Optional[T] {
	if o.IsPresent() && filter(o.Get()) {
		return Of(o.Get())
	}
	return Empty[T]()
}

// MapToAny applies the provided mapper function to the value contained in the Optional if present,
// returning a new Optional containing the mapped value; otherwise, it returns an empty Optional.
func (o Optional[T]) MapToAny(mapper func(T) any) Optional[any] {
	if o.IsPresent() {
		return Of(mapper(o.Get()))
	}
	return Empty[any]()
}

// MapToString applies the provided mapper function to the value contained in the Optional if present,
// returning a new Optional containing the mapped string value; otherwise, it returns an empty Optional.
func (o Optional[T]) MapToString(mapper func(T) string) Optional[string] {
	if o.IsPresent() {
		return Of(mapper(o.Get()))
	}
	return Empty[string]()
}

// MapToInt applies the provided mapper function to the value contained in the Optional if present,
// returning a new Optional containing the mapped int value; otherwise, it returns an empty Optional.
func (o Optional[T]) MapToInt(mapper func(T) int) Optional[int] {
	if o.IsPresent() {
		return Of(mapper(o.Get()))
	}
	return Empty[int]()
}

// IsEmpty returns true if the Optional does not contain a value; otherwise, it returns false.
func (o Optional[T]) IsEmpty() bool {
	return o.value == nil
}

// IsPresent returns true if the Optional contains a value; otherwise, it returns false.
func (o Optional[T]) IsPresent() bool {
	return o.value != nil
}

// IfPresent executes the provided consumer function with the value contained in the Optional if present.
func (o Optional[T]) IfPresent(consumer func(T)) {
	if o.IsPresent() {
		consumer(o.Get())
	}
}

// IfPresentOrElse executes the provided consumer function with the value if present; otherwise, it executes the provided alternative function.
func (o Optional[T]) IfPresentOrElse(consumer func(T), or func()) {
	if o.IsPresent() {
		consumer(o.Get())
	}
	or()
}

// OrElseGet returns the value contained in the Optional if present; otherwise, it invokes the provided supplier function to obtain the value.
func (o Optional[T]) OrElseGet(supplier func() T) T {
	if o.IsPresent() {
		return o.Get()
	}
	return supplier()
}

// Empty returns an empty Optional.
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

// Of creates an Optional containing the provided value.
func Of[T any](value T) Optional[T] {
	return OfPtr(&value)
}

// OfPtr creates an Optional from a pointer to a value.
// If the pointer is nil, it returns an empty Optional.
func OfPtr[T any](value *T) Optional[T] {
	return Optional[T]{value}
}

// FromTuplePtr takes a pointer value and error tuple
// creates an Optional containing the provided value if err is nil.
func FromTuplePtr[T any](value *T, err error) Optional[T] {
	if err != nil {
		return Empty[T]()
	}
	return OfPtr(value)
}

// FromTuple takes a value and error tuple
// creates an Optional containing the provided value if err is nil.
func FromTuple[T any](value T, err error) Optional[T] {
	return FromTuplePtr(&value, err)
}

// AndThenPtr takes a pointer value and error tuple
// creates an Optional containing the provided value if error is nil.
func (o Optional[T]) AndThenPtr(other *T, err error) Optional[T] {
	if o.IsPresent() && err == nil {
		return OfPtr(other)
	}
	return Empty[T]()
}

// AndThen takes a value and error tuple
// creates an Optional containing the provided value if error is nil.
func (o Optional[T]) AndThen(other T, err error) Optional[T] {
	return o.AndThenPtr(&other, err)
}

// Map applies the provided mapper function to the value contained in the Optional if present,
// returning a new Optional containing the mapped value; otherwise, it returns an empty Optional.
func Map[T, R any](o Optional[T], mapper func(T) R) Optional[R] {
	if o.IsPresent() {
		return Of(mapper(o.Get()))
	}
	return Empty[R]()
}
