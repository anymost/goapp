package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func (person *Person) sayName() {
	fmt.Println(person.name)
}

func (person *Person) sayAge() {
	fmt.Println(person.age)
}

func (person *Person) String() string {
	return person.name + strconv.Itoa(person.age)
}

type NameSayer interface {
	sayName()
}

func sayName(sayer NameSayer) {
	sayer.sayName()
}
