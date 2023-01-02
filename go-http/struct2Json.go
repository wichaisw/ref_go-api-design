package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	v := User{
		Id: 1, Name: "Kradas", Age: 12,
	}

	// struct to JSON
	b, err := json.Marshal(v)

	fmt.Printf("value: %v \n", b)
	fmt.Printf("type: %T \n", b)
	fmt.Printf("string: %s \n", b)
	fmt.Printf("Error: %v \n", err)
}
