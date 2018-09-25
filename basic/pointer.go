package main

import "fmt"

func basicPointer()  {
	value := 1
	var pointer *int
	pointer = &value
	fmt.Println(pointer)
	fmt.Println(*pointer)
}

func equalPointer()  {
	value := 1
	fmt.Println(value == *(&value))
}

func changeValueFromPointer()  {
	value := "hello"
	var pointer *string
	pointer = &value
	*pointer = "world"
	fmt.Println(value)
}

func changeArgsValue(value *int)  {
	*value = 2
}



func main() {
	// basicPointer()
	// equalPointer()
	// changeValueFromPointer()
	//value := 1
	//changeArgsValue(&value)
	//fmt.Println(value)
	a := 1
	var b *int
	b = &a
	*b = 2
	fmt.Println(a)
}
