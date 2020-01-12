package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {

	const json = `{
	"name": {
		"first": "Amitesh",
		"last": "Kumar"
	},
	"age": 30
}`

	firstName := gjson.Get(json, "name.first")
	lastName := gjson.Get(json, "name.last")
	age := gson.Get(json, "age")
	fmt.Printf("User Details -> First Name: %v, Last Name: %v, Age: %d", firstName, lastName, age)
}
