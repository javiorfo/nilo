# nilo
*Go Optionals library for handling nil values and some errors (partially inspired in Java Optionals)*

## Caveats
- This plugin requires Go 1.23

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

  nilo.FromTuplePtr(getUser(true)).AndThen(getUser2(true)).IfPresent(print)

  _, err := test(false).OrError(errors.New("some err"))
  fmt.Println(err.Error())

  fmt.Println(test(true).MapToString(func(v string) string { return v + ", World" }).OrElse("another string"))
}
```

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)
