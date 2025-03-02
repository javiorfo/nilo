package main

import (
	"errors"
	"fmt"
	"slices"

	"github.com/javiorfo/nilo"
)

func main() {
	var numbers []int

	evenNumbersOpt := nilo.Of(numbers).Filter(func(n []int) bool {
		return slices.Contains(n, 1)
	})

	if evenNumbersOpt.IsPresent() {
		fmt.Println("One is present")
	}

	nilo.FromTuplePtr(getUser(true)).AndThen(getUser2(true)).IfPresent(print)

	_, err := test(false).OrError(errors.New("some err"))
	fmt.Println(err.Error())

	fmt.Println(test(true).MapToString(func(v string) string { return v + ", World" }).OrElse("another string"))
}

func test(b bool) nilo.Optional[string] {
	if b {
		return nilo.Of("Hello")
	}
	return nilo.Empty[string]()
}

func print[T any](v T) {
	fmt.Printf("Value: %#v\n", v)
}

func getUser(b bool) (*int, error) {
	if b {
		i := 1
		return &i, nil
	}
	return nil, errors.New("error")
}

func getUser2(b bool) (int, error) {
	if b {
		return 2, nil
	}
	return 0, errors.New("error")
}
