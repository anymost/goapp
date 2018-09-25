package main

import (
	"fmt"
	"runtime"
)

func ifPractice(val int)  {
	if val == 1 {
		fmt.Println("a is 1")
	} else if val == 2 {
		fmt.Println("a is 2")
	} else if val == 3 {
		fmt.Println("a is 3")
	} else {
		fmt.Println("a is not 1,2,3")
	}

	fmt.Println(runtime.GOOS)
}

func ifInit()  {
	val := 0
	// if中会有单独的块级作用域
	if val = 1; val > 0 {
		fmt.Println("a ")
	}
	fmt.Println(val)
}

func isError(val int) (int, bool)  {
	if val > 0 {
		return val, true
	}
	return 0, false
}

func main()  {
	ifInit()
}