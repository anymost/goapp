package main

import (
	"fmt"
)

func basicSwitch()  {
	const name = 1
	switch name {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println(4)
	}
}

func multiSwitch()  {
	const name = 1
	switch name {
	case 1, 2:
		fmt.Println("is 1~2")
	case 3, 4:
		fmt.Println("is 3~4")
	default:
		fmt.Println("is > 4")
	}
}

func ifSwitch() {
	name := 1
	switch  {
	case name < 1:
		fmt.Println("name less 1")
	case name == 1:
		fmt.Println("name equal 1")
	case name > 1:
		fmt.Println("name larger 1")
	}
}

func initSwitch()  {
	switch name := 1; {
	case name == 1:
		fmt.Println("name equal 1")
	default:
		fmt.Println("name not equal 1")
	}
}



func main() {
	// basicSwitch()
	// multiSwitch()
	// ifSwitch()
	// initSwitch()
}
