package main

import (
	"encoding/json"
	"fmt"
)

// capital letter field is public (could access & mutate from other package)
type User struct {
	Id   int
	Name string `json:"nickname"`
	Age  int
}

func main() {
	// name value won't be assigned since we have nickname as its tag
	// GO will look for struct tag if provided
	data := []byte(`{"id": 1, "nickname": "Kradas", "age": 12}`)

	// need pointer because User is in package main, while Unmarshal is in package json
	// so we sent User{} reference so that Unmarshal could put value in variable u
	u := &User{}
	fmt.Printf("before unmarshal value: %#v\n", u)
	err := json.Unmarshal(data, u)

	// if declare as struct value (not struct pointer), need to pass reference using &u
	// var u User
	// json.Unmarshal(data, &u)

	fmt.Printf("struct value: %v\n", u)
	fmt.Printf("struct value with field name: %+v\n", u)
	fmt.Printf("Go syntax representation of the value: %#v\n", u)
	fmt.Println(err)
}
