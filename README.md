# nilo
*Go Option library for handling nil values, some errors and JSON marshaling*

## Caveats
- This library requires Go 1.23+

## Installation
```bash
go get -u github.com/javiorfo/nilo@latest
```

## Example
#### Examples [here](https://github.com/javiorfo/nilo/tree/master/examples)
```go
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
  var optUser = nilo.Nil[User]()

  fmt.Printf("User or default: %+v\n", optUser.OrDefault())
  fmt.Printf("User or: %+v\n", optUser.Or(*new(User)))
  fmt.Printf("User or else: %+v\n", optUser.OrElse(func() User { return User{"else"} }))
  fmt.Printf("Map or: %+v\n", optUser.Map(func(u User) User {
	  u.Name = "something"
	  return u
  }).Or(User{"or"}))

  nilo.FromResult(getUser(true)).
    AndResult(getUser2).
    Consume(print)

  _, err := test(false).OrError(func() error { return errors.New("some err") })
  fmt.Println("Error:", err.Error())

  fmt.Println(test(true).
    MapToString(func(v string) string { return v + ", World" }).
    Or("another string"))
}

func test(b bool) nilo.Option[string] {
  if b {
	  return nilo.Value("Hello")
  }
  return nilo.Nil[string]()
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
```

#### JSON marshal
```go
package main

import (
  "encoding/json"
  "fmt"

  "github.com/javiorfo/nilo"
)

type User struct {
  Name string              `json:"name"`
  Code nilo.Option[string] `json:"code"`
}

func main() {
  var unmarshalUser User
  user := User{
	  Name: "Name",
	  Code: nilo.Nil[string](),
  }

  // Marshal
  jsonData, err := json.MarshalIndent(user, "", "  ")
  if err != nil {
	  fmt.Println("Error marshaling to JSON:", err)
	  return
  }

  fmt.Println(string(jsonData))
	
  // Unmarshal
  err = json.Unmarshal(jsonData, &unmarshalUser)
  if err != nil {
	  fmt.Println("Error unmarshaling:", err)
	  return
  }
	
  fmt.Printf("Unmarshaled User: %+v\n", unmarshalUser)
  if unmarshalUser.Code.IsNil() {
	  fmt.Printf("Code is Nil: %s\n", unmarshalUser.Code)
  }

  // Put Some in Code
  user.Code.Insert("code")

  // Marshal
  jsonData, err = json.MarshalIndent(user, "", "  ")
  if err != nil {
	  fmt.Println("Error marshaling to JSON:", err)
	  return
  }

  fmt.Println(string(jsonData))

  // Unmarshal
  err = json.Unmarshal(jsonData, &unmarshalUser)
  if err != nil {
	  fmt.Println("Error unmarshaling:", err)
	  return
  }
	
  fmt.Printf("Unmarshaled User: %+v\n", unmarshalUser)
  if unmarshalUser.Code.IsValue() {
	  fmt.Printf("Code is Value: %s\n", unmarshalUser.Code)
  }
}
```

#### All methods and functions
```go
func (o Option[T]) AsValue() T
func (o Option[T]) AsPtr() *T
func (o Option[T]) Or(other T) T
func (o Option[T]) OrDefault() T
func (o Option[T]) OrElse(supplier func() T) T
func (o Option[T]) OrError(err func() error) (*T, error)
func (o Option[T]) OrPanic(msg string) T
func (o Option[T]) Filter(filter func(T) bool) Option[T]
func (o Option[T]) IsNil() bool
func (o Option[T]) IsValue() bool
func (o Option[T]) IsValueAnd(predicate func(T) bool) bool
func (o Option[T]) IsNilOr(predicate func(T) bool) bool
func (o Option[T]) IfNil(executor func())
func (o Option[T]) Inspect(inspector func(T)) Option[T]
func (o Option[T]) Consume(consumer func(T))
func (o *Option[T]) Take() Option[T]
func (o *Option[T]) TakeIf(predicate func(T) bool) Option[T]
func (o *Option[T]) Insert(value T)
func (o Option[T]) Map(mapper func(T) T) Option[T]
func (o Option[T]) MapToString(mapper func(T) string) Option[string]
func (o Option[T]) MapToInt(mapper func(T) int) Option[int]
func (o Option[T]) MapToBool(mapper func(T) bool) Option[bool]
func (o Option[T]) MapOrDefault(mapper func(T) T) T
func (o Option[T]) AndThen(fn func(T) Option[T]) Option[T]
func (o Option[T]) AndOk(apply func(T) (T, error)) Option[T]
func (o Option[T]) AndOkPtr(apply func(T) (*T, error)) Option[T]
func (o Option[T]) MarshalJSON() ([]byte, error)
func (o *Option[T]) UnmarshalJSON(data []byte) error
func (o Option[T]) String() string
func (o Option[T]) Iter() iter.Seq[T] {
func Ok[T any](value T, err error) Option[T]
func Nil[T any]() Option[T]
func Value[T any](value T) Option[T]
func Ptr[T any](value *T) Option[T]
func Cast[T any](value T) Option[T]
```

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
