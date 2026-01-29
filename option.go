// Package nilo provides a generic Option type that can be used to represent a value that may or may not be present.
package nilo

// Option is a generic type that represents an option value.
// An `Option` can either be `Value`, containing a value of type `T`, or `Nil`,
// indicating the absence of a value.
type Option[T any] struct {
	value *T
}

// AsValue returns the contained value of an `Option`.
//
// Panics if the `Option` is `Nil`. This method should only be used when
// you are certain the `Option` contains a value.
func (o Option[T]) AsValue() T {
	if o.value == nil {
		panic("Option value is Nil")
	}
	return *o.value
}

// Or returns the contained value if the `Option` is `Value`, otherwise
// returns the provided default value `other`.
//
// This is a safe alternative to `AsValue` when you have a default value to use.
//
// Parameters:
//   - other: The default value to return if the `Option` is `Nil`.
func (o Option[T]) Or(other T) T {
	if o.IsValue() {
		return o.AsValue()
	}
	return other
}

// AsPtr returns the pointer of the contained value without checking
// if the `Option` is `Value`.
//
// The caller is responsible for ensuring the `Option` is  `Value` before calling
// this method. Calling this on a `Nil` `Option` will result in a nil pointer.
func (o Option[T]) AsPtr() *T {
	return o.value
}

// OrDefault returns the contained value if the `Option` is `Value`.
// If the `Option` is `Nil`, it returns a default value.
//
// The default value is determined by the following:
//  1. If the type `T` implements the `Default` interface, `Default()` is called
//     to get the default value.
//  2. Otherwise, the Go language's zero value for type `T` is returned.
func (o Option[T]) OrDefault() T {
	if o.IsValue() {
		return o.AsValue()
	}

	return defaultImplOrNew[T]()
}

// OrElse returns the contained value if the `Option` is `Value`,
// otherwise it calls the provided `supplier` function to get a default value.
//
// This is useful for providing a default value that is expensive to compute,
// as the supplier function is only called when needed.
//
// Parameters:
//   - supplier: A function that returns the default value if the `Option` is `Nil`.
func (o Option[T]) OrElse(supplier func() T) T {
	if o.IsValue() {
		return o.AsValue()
	}
	return supplier()
}

// OrError converts an `Option` into a `(value, error)` tuple.
//
// If the `Option` is `Value`, it returns a pointer to the value and a `nil` error.
// If the `Option` is `Nil`, it returns a `nil` pointer and the error returned
// by the `err` supplier function. This is useful when creating the error is
// an expensive operation, as the supplier function is only called when needed.
//
// Parameters:
//   - err: A function that returns the error to be used if the `Option` is `Nil`.
func (o Option[T]) OrError(err func() error) (*T, error) {
	if o.IsValue() {
		return o.value, nil
	}
	return nil, err()
}

// Filter calls a predicate function on the contained value if the `Option` is `Value`.
//
// If the `Option` is `Value` and the predicate returns `true`, it returns a
// `Value` `Option` with the original value. Otherwise, it returns a `Nil` `Option`.
// This is useful for removing values that do not meet a certain condition.
//
// Parameters:
//   - filter: A predicate function that returns `true` or `false` based on the value.
func (o Option[T]) Filter(filter func(T) bool) Option[T] {
	if o.IsValue() && filter(o.AsValue()) {
		return Value(o.AsValue())
	}
	return Nil[T]()
}

// IsNil returns `true` if the `Option` is `Nil`.
func (o Option[T]) IsNil() bool {
	return o.value == nil
}

// IsValue returns `true` if the `Option` is `Value`.
func (o Option[T]) IsValue() bool {
	return o.value != nil
}

// Inspect calls a function on the contained value if the `Option` is `Value`,
// and then returns the original `Option`.
//
// This is useful for debugging or logging the value without consuming the `Option`.
//
// Parameters:
//   - consumer: A function that takes the `Option`'s value.
func (o Option[T]) Inspect(consumer func(T)) Option[T] {
	if o.IsValue() {
		consumer(o.AsValue())
	}
	return o
}

// Consume calls a function on the contained value if the `Option` is `Value`,
// and returns nothing.
//
// This is often used for side effects, such as updating state,
// without needing to manually check the presence of the value.
//
// Parameters:
//   - consumer: A function that takes the `Option`'s value.
func (o Option[T]) Consume(consumer func(T)) {
	if o.IsValue() {
		consumer(o.AsValue())
	}
}

// IfNil calls a function if the `Option` is `Nil`,
// and returns nothing.
//
// Parameters:
//   - executor: A function that takes the no arguments..
func (o Option[T]) IfNil(executor func()) {
	if o.IsNil() {
		executor()
	}
}

// Nil returns an empty Option.
func Nil[T any]() Option[T] {
	return Option[T]{}
}

// Value creates an Option containing the provided value.
func Value[T any](value T) Option[T] {
	return Option[T]{&value}
}

// Ptr creates an Option from a pointer to a value.
// If the pointer is nil, it returns an empty Option.
func Ptr[T any](value *T) Option[T] {
	return Option[T]{value}
}

// AndThen is a chaining method that applies a function to the contained value
// if the `Option` is `Value`, returning the result.
//
// If the `Option` is `Value`, the `fn` is called with the unwrapped value,
// and the new `Option` returned by `fn` is the result. This allows for
// a sequence of fallible operations. If the `Option` is `Nil`, this method
// returns `Nil` without calling `fn`.
//
// Parameters:
//   - fn: A function that takes the `Option`'s value and returns a new `Option`.
func (o Option[T]) AndThen(fn func(T) Option[T]) Option[T] {
	if o.IsValue() {
		return fn(o.AsValue())
	}
	return Nil[T]()
}

// IsValueAnd returns `true` if the `Option` is `Value` and the value contained
// within it satisfies the given `predicate`.
//
// If the `Option` is `Nil`, this method short-circuits and returns `false`.
// This is a convenient way to check for both presence and a condition in one call.
//
// Parameters:
//   - predicate: A function that takes the `Option`'s value and returns a boolean.
func (o Option[T]) IsValueAnd(predicate func(T) bool) bool {
	if o.IsValue() && predicate(o.AsValue()) {
		return true
	}
	return false
}

// IsNilOr returns `true` if the `Option` is `Nil` or if the contained value
// satisfies the given `predicate`.
//
// This acts as a logical "OR" on the state of the `Option`. If the `Option`
// is `Nil`, the predicate is not evaluated, and `true` is returned. If it
// is `Value`, the function returns the result of the `predicate`.
//
// Parameters:
//   - predicate: A function that takes the `Option`'s value and returns a boolean.
func (o Option[T]) IsNilOr(predicate func(T) bool) bool {
	if o.IsValue() {
		return predicate(o.AsValue())
	}
	return true
}

// Take takes the value out of the `Option`, leaving a `Nil` in its place.
//
// If the `Option` is `Value`, this method returns the original `Value` `Option`
// and sets the receiver to `Nil`. If the `Option` is `Nil`, it remains `Nil`.
func (o *Option[T]) Take() Option[T] {
	oldValue := *o
	o.value = nil
	return oldValue
}

// TakeIf takes the value out of the `Option` and leaves a `Nil` if the
// contained value satisfies the given `predicate`.
//
// If the `Option` is `Value` and the `predicate` returns `true`, it behaves
// identically to `Take`, returning the original `Value` `Option` and setting
// the receiver to `Nil`. Otherwise, it returns `Nil` and does not modify
// the original `Option`.
//
// Parameters:
//   - predicate: A function that tests the `Option`'s value.
func (o *Option[T]) TakeIf(predicate func(T) bool) Option[T] {
	if o.MapToBool(predicate).Or(false) {
		return o.Take()
	}
	return Nil[T]()
}

// OrPanic returns the contained `Value` value, but panics with a custom message
// if the `Option` is `Nil`.
//
// This method is used when a `Nil` value is considered an unrecoverable error
// in your program, and you want to crash with a descriptive message.
//
// Parameters:
//   - msg: The message to be used in the panic if the `Option` is `Nil`.
func (o Option[T]) OrPanic(msg string) T {
	if o.IsValue() {
		return o.AsValue()
	}
	panic(msg)
}

// Insert replaces the contained value with a new one.
//
// If the `Option` is currently `Nil`, it is changed to `Value` with the new value.
// If the `Option` is already `Value`, its existing value is replaced by the new one.
//
// Parameters:
//   - value: The new value to be inserted into the `Option`.
func (o *Option[T]) Insert(value T) {
	*o = Value(value)
}

// Cast attempts to assert the value any to type T.
// If the type assertion is successful, it returns an Option containing the value.
// If the assertion fails (e.g., incompatible types or i is nil), it returns a Nil Option.
//
// Example:
//
//	opt := Cast[int](anyValue)
func Cast[T any](value any) Option[T] {
	v, ok := value.(T)
	if ok {
		return Value(v)
	}
	return Nil[T]()
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
