# nilo
*Rusty Go Option library for handling nil values, some errors and marshaling*

## Caveats
- This library requires Go 1.23+

## Intallation
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
  var optUser = nilo.None[User]()

  fmt.Printf("User or default: %+v\n", optUser.UnwrapOrDefault())
  fmt.Printf("User or: %+v\n", optUser.UnwrapOr(*new(User)))
  fmt.Printf("User or else: %+v\n", optUser.UnwrapOrElse(func() User { return User{"else"} }))
  fmt.Printf("Map or: %+v\n", optUser.MapOr(User{"or"}, func(u User) User {
	  u.Name = "something"
	  return u
  }))

  nilo.FromResult(getUser(true)).OkAndResult(getUser2).Inspect(print)

  _, err := test(false).OkOr(errors.New("some err"))
  fmt.Println("Error:", err.Error())

  fmt.Println(test(true).MapToString(func(v string) string { return v + ", World" }).UnwrapOr("another string"))
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
	  Code: nilo.None[string](),
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
  if unmarshalUser.Code.IsNone() {
	  fmt.Printf("Code is None: %s\n", unmarshalUser.Code)
  }

  // Put Some in Code
  user.Code.Replace("some code")

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
  if unmarshalUser.Code.IsSome() {
	  fmt.Printf("Code is Some with value: %s\n", unmarshalUser.Code.Unwrap())
  }
}
```

#### All methods
```go
func (o Option[T]) Unwrap() T
func (o Option[T]) UnwrapOr(other T) T
func (o Option[T]) UnwrapUnchecked() *T
func (o Option[T]) UnwrapOrDefault() T
func (o Option[T]) UnwrapOrElse(supplier func() T) T
func (o Option[T]) OkOr(err error) (*T, error)
func (o Option[T]) OkOrElse(err func() error) (*T, error)
func (o Option[T]) OrElse(supplier func() Option[T]) Option[T]
func (o Option[T]) Filter(filter func(T) bool) Option[T]
func (o Option[T]) IsNone() bool
func (o Option[T]) IsSome() bool
func (o Option[T]) Inspect(consumer func(T)) Option[T]
func (o Option[T]) InspectOrElse(consumer func(T), or func())
func None[T any]() Option[T]
func Some[T any](value T) Option[T]
func SomePtr[T any](value *T) Option[T]
func (o Option[T]) AndThen(fn func(T) Option[T]) Option[T]
func (o Option[T]) And(other Option[T]) Option[T]
func (o Option[T]) Or(other Option[T]) Option[T]
func (o Option[T]) Xor(other Option[T]) Option[T]
func (o Option[T]) IsSomeAnd(predicate func(T) bool) bool
func (o Option[T]) IsNoneOr(predicate func(T) bool) bool
func (o Option[T]) Expect(msg string) T
func (o *Option[T]) Take() Option[T]
func (o *Option[T]) TakeIf(predicate func(T) bool) Option[T]
func (o *Option[T]) Replace(value T) Option[T]
func (o *Option[T]) Insert(value T)
func (o *Option[T]) GetOrInsert(value T) T
func (o *Option[T]) GetOrInsertWith(supplier func() T)
func (o *Option[T]) GetOrInsertDefault() T
func (o Option[T]) Map(mapper func(T) T) Option[T]
func (o Option[T]) MapToString(mapper func(T) string) Option[string]
func (o Option[T]) MapToInt(mapper func(T) int) Option[int]
func (o Option[T]) MapToBool(mapper func(T) bool) Option[bool]
func (o Option[T]) MapOr(def T, mapper func(T) T) T
func (o Option[T]) MapOrString(def string, mapper func(T) string) string
func (o Option[T]) MapOrInt(def int, mapper func(T) int) int
func (o Option[T]) MapOrBool(def bool, mapper func(T) bool) bool
func (o Option[T]) MapOrElse(supplier func() T, mapper func(T) T) T
func (o Option[T]) MapOrElseString(supplier func() string, mapper func(T) string) string
func (o Option[T]) MapOrElseInt(supplier func() int, mapper func(T) int) int
func (o Option[T]) MapOrElseBool(supplier func() bool, mapper func(T) bool) bool
func (o Option[T]) MapOrDefault(mapper func(T) T) T
func (o Option[T]) OkAndResult(apply func(T) (T, error)) Option[T]
func FromResult[T any](value T, err error) Option[T]
func (o Option[T]) MarshalJSON() ([]byte, error)
func (o *Option[T]) UnmarshalJSON(data []byte) error
func (o Option[T]) String() string
```

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
