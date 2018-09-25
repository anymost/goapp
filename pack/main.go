package main

import (
	"fmt"
)

func main() {
	person := Person{"hello", 1}
	person.sayName()
	sayName(&person)
	fmt.Println(person)
}
