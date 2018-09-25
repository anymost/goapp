package main

import (
	"fmt"
	"strconv"
)

func commonFor()  {
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}
}

func layerFor()  {
	for i := 0; i < 10; i++ {
		for a:= 0; a < 10; a++ {
			fmt.Println(i + a)
		}
	}
}

func whileFor()  {
	a := 0
	for a < 10 {
		fmt.Println(a)
		a++
	}
}

func infinityFor()  {
	for {
		fmt.Println("hello")
	}
}

func rangeFor()  {
	array := []int{1, 2, 3, 4}
	for index, value := range array {
		fmt.Println("index is " + strconv.Itoa(index) + "value is " + strconv.Itoa(value))
	}
}

func main() {
	// commonFor()
	// layerFor()
	// whileFor()
	// infinityFor()
	rangeFor()
}
