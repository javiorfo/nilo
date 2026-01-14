// Package nilo provides a generic Option type that can be used to represent a value that may or may not be present.
package nilo

// Option is a generic type that represents an option value.
// An `Option` can either be `Some`, containing a value of type `T`, or `None`,
// indicating the absence of a value.
type Option[T any] struct {
	value *T
}

// Unwrap returns the contained value of a `Some` `Option`.
//
// Panics if the `Option` is `None`. This method should only be used when
// you are certain the `Option` contains a value.
func (o Option[T]) Unwrap() T {
	return *o.value
}

// UnwrapOr returns the contained value if the `Option` is `Some`, otherwise
// returns the provided default value `other`.
//
// This is a safe alternative to `Unwrap` when you have a default value to use.
//
// Parameters:
//   - other: The default value to return if the `Option` is `None`.
func (o Option[T]) UnwrapOr(other T) T {
	if o.IsSome() {
		return o.Unwrap()
	}
	return other
}

// UnwrapUnchecked returns a pointer to the contained value without checking
// if the `Option` is `Some`.
//
// The caller is responsible for ensuring the `Option` is `Some` before calling
// this method. Calling this on a `None` `Option` will result in a nil pointer.
func (o Option[T]) UnwrapUnchecked() *T {
	return o.value
}

// UnwrapOrDefault returns the contained value if the `Option` is `Some`.
// If the `Option` is `None`, it returns a default value.
//
// The default value is determined by the following:
//  1. If the type `T` implements the `Default` interface, `Default()` is called
//     to get the default value.
//  2. Otherwise, the Go language's zero value for type `T` is returned.
func (o Option[T]) UnwrapOrDefault() T {
	if o.IsSome() {
		return o.Unwrap()
	}

	return defaultImplOrNew[T]()
}

// UnwrapOrElse returns the contained value if the `Option` is `Some`,
// otherwise it calls the provided `supplier` function to get a default value.
//
// This is useful for providing a default value that is expensive to compute,
// as the supplier function is only called when needed.
//
// Parameters:
//   - supplier: A function that returns the default value if the `Option` is `None`.
func (o Option[T]) UnwrapOrElse(supplier func() T) T {
	if o.IsSome() {
		return o.Unwrap()
	}
	return supplier()
}

// OkOr converts an `Option` into a Go-style `(value, error)` tuple.
//
// If the `Option` is `Some`, it returns a pointer to the value and a `nil` error.
// If the `Option` is `None`, it returns a `nil` pointer and the provided `err`.
// This is useful for integrating `Option` with functions that return an error.
//
// Parameters:
//   - err: The error to return if the `Option` is `None`.
func (o Option[T]) OkOr(err error) (*T, error) {
	if o.IsSome() {
		return o.value, nil
	}
	return nil, err
}

// OkOrElse converts an `Option` into a `(value, error)` tuple.
//
// If the `Option` is `Some`, it returns a pointer to the value and a `nil` error.
// If the `Option` is `None`, it returns a `nil` pointer and the error returned
// by the `err` supplier function. This is useful when creating the error is
// an expensive operation, as the supplier function is only called when needed.
//
// Parameters:
//   - err: A function that returns the error to be used if the `Option` is `None`.
func (o Option[T]) OkOrElse(err func() error) (*T, error) {
	if o.IsSome() {
		return o.value, nil
	}
	return nil, err()
}

// OrElse returns the `Option` itself if it is `Some`.
// If the `Option` is `None`, it calls the `supplier` function and returns
// the resulting `Option`.
//
// This is useful for providing a fallback `Option` to use if the original
// is empty. The `supplier` function is only called if the `Option` is `None`,
// which can be beneficial for performance if creating the fallback is an
// expensive operation.
//
// Parameters:
//   - supplier: A function that provides the fallback `Option` to return if
//     the original `Option` is `None`.
func (o Option[T]) OrElse(supplier func() Option[T]) Option[T] {
	if o.IsNone() {
		return supplier()
	}
	return Some(o.Unwrap())
}

// Filter calls a predicate function on the contained value if the `Option` is `Some`.
//
// If the `Option` is `Some` and the predicate returns `true`, it returns a
// `Some` `Option` with the original value. Otherwise, it returns a `None` `Option`.
// This is useful for removing values that do not meet a certain condition.
//
// Parameters:
//   - filter: A predicate function that returns `true` or `false` based on the value.
func (o Option[T]) Filter(filter func(T) bool) Option[T] {
	if o.IsSome() && filter(o.Unwrap()) {
		return Some(o.Unwrap())
	}
	return None[T]()
}

// IsNone returns `true` if the `Option` is `None`.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// IsSome returns `true` if the `Option` is `Some`.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// Inspect calls a function on the contained value if the `Option` is `Some`,
// and then returns the original `Option`.
//
// This is useful for debugging or logging the value without consuming the `Option`.
//
// Parameters:
//   - consumer: A function that takes the `Option`'s value.
func (o Option[T]) Inspect(consumer func(T)) Option[T] {
	if o.IsSome() {
		consumer(o.Unwrap())
	}
	return o
}

// InspectOrElse calls one of two functions depending on whether the `Option` is `Some` or `None`.
//
// If the `Option` is `Some`, it calls the `consumer` function with the contained value.
// If the `Option` is `None`, it calls the `or` function.
// This is useful for debugging or performing side effects based on the `Option`'s state.
//
// Parameters:
//   - consumer: A function that takes the `Option`'s value, called if the `Option` is `Some`.
//   - or: A function with no parameters, called if the `Option` is `None`.
func (o Option[T]) InspectOrElse(consumer func(T), or func()) {
	if o.IsSome() {
		consumer(o.Unwrap())
	} else {
		or()
	}
}

// Consume calls a function on the contained value if the `Option` is `Some`,
// and returns nothing.
//
// This is often used for side effects, such as updating state,
// without needing to manually check the presence of the value.
//
// Parameters:
//   - consumer: A function that takes the `Option`'s value.
func (o Option[T]) Consume(consumer func(T)) {
	if o.IsSome() {
		consumer(o.Unwrap())
	}
}

// None returns an empty Option.
func None[T any]() Option[T] {
	return Option[T]{}
}

// Some creates an Option containing the provided value.
func Some[T any](value T) Option[T] {
	return Option[T]{&value}
}

// SomePtr creates an Option from a pointer to a value.
// If the pointer is nil, it returns an empty Option.
func SomePtr[T any](value *T) Option[T] {
	return Option[T]{value}
}

// AndThen is a chaining method that applies a function to the contained value
// if the `Option` is `Some`, returning the result.
//
// If the `Option` is `Some`, the `fn` is called with the unwrapped value,
// and the new `Option` returned by `fn` is the result. This allows for
// a sequence of fallible operations. If the `Option` is `None`, this method
// returns `None` without calling `fn`.
//
// Parameters:
//   - fn: A function that takes the `Option`'s value and returns a new `Option`.
func (o Option[T]) AndThen(fn func(T) Option[T]) Option[T] {
	if o.IsSome() {
		return fn(o.Unwrap())
	}
	return None[T]()
}

// And combines two `Option`s.
//
// If both the current `Option` and the `other` `Option` are `Some`, it returns
// the `other` `Option`. In all other cases (if either `Option` is `None`), it
// returns `None`. This can be used to perform a logical "AND" operation on
// `Option` values.
//
// Parameters:
//   - other: The other `Option` to combine with the current one.
func (o Option[T]) And(other Option[T]) Option[T] {
	if o.IsSome() && other.IsSome() {
		return other
	}
	return None[T]()
}

// Or returns the `Option` if it is `Some`, otherwise it returns the `other` `Option`.
//
// This method acts as a logical "OR" operation. If the current `Option` is `Some`,
// its value is preferred. If it's `None`, it falls back to the `other` `Option`.
//
// Parameters:
//   - other: The alternative `Option` to return if the current one is `None`.
func (o Option[T]) Or(other Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}

	if other.IsSome() {
		return other
	}

	return None[T]()
}

// Xor returns the `Option` that is `Some`, if exactly one of them is `Some`.
//
// This performs a logical "XOR" (exclusive OR) operation. It returns a `Some`
// `Option` only if either the current `Option` is `Some` and the `other` is `None`,
// or vice versa. If both are `Some` or both are `None`, it returns `None`.
//
// Parameters:
//   - other: The other `Option` to perform the XOR operation with.
func (o Option[T]) Xor(other Option[T]) Option[T] {
	if o.IsSome() && other.IsNone() {
		return o
	}

	if other.IsSome() && o.IsNone() {
		return other
	}

	return None[T]()
}

// IsSomeAnd returns `true` if the `Option` is `Some` and the value contained
// within it satisfies the given `predicate`.
//
// If the `Option` is `None`, this method short-circuits and returns `false`.
// This is a convenient way to check for both presence and a condition in one call.
//
// Parameters:
//   - predicate: A function that takes the `Option`'s value and returns a boolean.
func (o Option[T]) IsSomeAnd(predicate func(T) bool) bool {
	if o.IsSome() && predicate(o.Unwrap()) {
		return true
	}
	return false
}

// IsNoneOr returns `true` if the `Option` is `None` or if the contained value
// satisfies the given `predicate`.
//
// This acts as a logical "OR" on the state of the `Option`. If the `Option`
// is `None`, the predicate is not evaluated, and `true` is returned. If it
// is `Some`, the function returns the result of the `predicate`.
//
// Parameters:
//   - predicate: A function that takes the `Option`'s value and returns a boolean.
func (o Option[T]) IsNoneOr(predicate func(T) bool) bool {
	if o.IsSome() {
		return predicate(o.Unwrap())
	}
	return true
}

// Take takes the value out of the `Option`, leaving a `None` in its place.
//
// If the `Option` is `Some`, this method returns the original `Some` `Option`
// and sets the receiver to `None`. If the `Option` is `None`, it remains `None`.
func (o *Option[T]) Take() Option[T] {
	oldValue := *o
	o.value = nil
	return oldValue
}

// TakeIf takes the value out of the `Option` and leaves a `None` if the
// contained value satisfies the given `predicate`.
//
// If the `Option` is `Some` and the `predicate` returns `true`, it behaves
// identically to `Take`, returning the original `Some` `Option` and setting
// the receiver to `None`. Otherwise, it returns `None` and does not modify
// the original `Option`.
//
// Parameters:
//   - predicate: A function that tests the `Option`'s value.
func (o *Option[T]) TakeIf(predicate func(T) bool) Option[T] {
	if o.MapOrBool(false, predicate) {
		return o.Take()
	}
	return None[T]()
}

// Replace replaces the contained value with a new value and returns the old `Option`.
//
// If the `Option` is `Some`, its value is updated to `value`. The original `Some`
// `Option` with the old value is returned. If the `Option` is `None`, it remains
// `None`, and a new `Some` `Option` with the new value is returned.
//
// Parameters:
//   - value: The new value to place into the `Option`.
func (o *Option[T]) Replace(value T) Option[T] {
	*o = Some(value)
	return *o
}

// Expect returns the contained `Some` value, but panics with a custom message
// if the `Option` is `None`.
//
// This method is used when a `None` value is considered an unrecoverable error
// in your program, and you want to crash with a descriptive message.
//
// Parameters:
//   - msg: The message to be used in the panic if the `Option` is `None`.
func (o Option[T]) Expect(msg string) T {
	if o.IsSome() {
		return o.Unwrap()
	}
	panic(msg)
}

// Insert replaces the contained value with a new one.
//
// If the `Option` is currently `None`, it is changed to `Some` with the new value.
// If the `Option` is already `Some`, its existing value is replaced by the new one.
//
// Parameters:
//   - value: The new value to be inserted into the `Option`.
func (o *Option[T]) Insert(value T) {
	*o = Some(value)
}

// GetOrInsert returns the contained value if the `Option` is `Some`. If the
// `Option` is `None`, it inserts the provided `value` into the `Option`,
// making it `Some`, and then returns the newly inserted value.
//
// Parameters:
//   - value: The value to insert if the `Option` is `None`.
func (o *Option[T]) GetOrInsert(value T) T {
	return o.GetOrInsertWith(func() T { return value })
}

// GetOrInsertWith returns the contained value if the `Option` is `Some`.
// If the `Option` is `None`, it calls the `supplier` function, inserts the
// returned value into the `Option`, and then returns the new value.
//
// This method is useful for providing a default value that is expensive to
// compute, as the `supplier` function is only called when needed.
//
// Parameters:
//   - supplier: A function that provides the value to be inserted if the
//     `Option` is `None`.
func (o *Option[T]) GetOrInsertWith(supplier func() T) T {
	if o.IsNone() {
		*o = Some(supplier())
	}
	return o.Unwrap()
}

// GetOrInsertDefault returns the contained value if the `Option` is `Some`.
// If the `Option` is `None`, it inserts a default value, making it `Some`,
// and then returns the new value.
//
// The default value is determined by the following:
//  1. If the type `T` implements the `Default` interface, its `Default()` method
//     is called to get the value.
//  2. Otherwise, the Go language's zero value for type `T` is used.
func (o *Option[T]) GetOrInsertDefault() T {
	return o.GetOrInsertWith(func() T {
		return defaultImplOrNew[T]()
	})
}

func defaultImplOrNew[T any]() T {
	var def T
	switch any(def).(type) {
	case Default[T]:
		return any(def).(Default[T]).Default()
	default:
		return *new(T)
	}
}
