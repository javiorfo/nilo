# nilo
*Go Optionals library for handling nil values and some errors (partially inspired in Java Optionals)*

## Caveats
- This library requires Go 1.23+

## Intallation
```bash
go get -u github.com/javiorfo/nilo@latest
```

## Example
#### Examples [here](https://github.com/javiorfo/nilo/tree/master/examples/example.go)
```go
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

  nilo.FromTuple(getUser(true)).AndThen(getUser2).IfPresent(print)

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

func getUser2(p *int) (*int, error) {
  if *p == 0 {
    return 0, errors.New("error")
  }
  r := *p + 2
  return &r, nil
}
```

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
