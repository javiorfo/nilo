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
	user.Code.Insert("some code")

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
		fmt.Printf("Code is Value: %s\n", unmarshalUser.Code.OrPanic("should not be nil"))
	}
}
