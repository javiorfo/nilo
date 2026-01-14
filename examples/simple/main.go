package main

import (
	"errors"
	"fmt"

	"github.com/javiorfo/nilo"
)

type User struct {
	Name string
}

// Implements nilo.Default interface
func (u User) Default() User {
	return User{"Default Implementation"}
}

func main() {
	var optUser = nilo.None[User]()

	fmt.Printf("User or default: %+v\n", optUser.UnwrapOrDefault())
	fmt.Printf("User or: %+v\n", optUser.UnwrapOr(*new(User)))
	fmt.Printf("User or else: %+v\n", optUser.UnwrapOrElse(func() User { return User{"else"} }))
	fmt.Printf("Map or: %+v\n", optUser.MapOr(User{"or"}, func(u User) User {
		u.Name = "something"
		return u
	}))

	nilo.FromResult(getUser(true)).
		OkAndResult(getUser2).
		Consume(print)

	_, err := test(false).OkOr(errors.New("some err"))
	fmt.Println("Error:", err.Error())

	fmt.Println(test(true).
		MapToString(func(v string) string { return v + ", World" }).
		UnwrapOr("another string"))
}

func test(b bool) nilo.Option[string] {
	if b {
		return nilo.Some("Hello")
	}
	return nilo.None[string]()
}

func print[T any](v *T) {
	fmt.Printf("Value: %#v\n", *v)
}

func getUser(b bool) (*int, error) {
	if b {
		i := 1
		return &i, nil
	}
	return nil, errors.New("error")
}

func getUser2(v *int) (*int, error) {
	if *v == 0 {
		return nil, errors.New("error")
	}
	i := *v + 2
	return &i, nil
}
