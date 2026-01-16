package nilo

// Default is an interface that can be implemented by types that wish to provide
// a custom 'default' value. This is used by methods on `Option` such as
// `OrDefault` or `MapOrDefault` to create an instance of the type
// when a `Nil` `Option` is encountered.
//
// Types that implement this interface can define their own logic for what
// constitutes a 'default' value, instead of relying on the Go language's
// zero value.
type Default[T any] interface {
	Default() T
}

// Simple function to wrap in OrError method
func ReturnError(err error) func() error {
	return func() error {
		return err
	}
}
