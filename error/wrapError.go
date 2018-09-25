package main

import "fmt"

/**
嵌套的panic 会保证所有的defer都能执行 但是defer要先于panic定义
 */

func cPanic()  {
	defer func() {
		fmt.Println("c panic execute")
	}()
	fmt.Println("start")
	panic("died")
	fmt.Println("end")

}

func bPanic()  {
	defer func() {
		fmt.Println("b panic execute")
	}()
	cPanic()
}

func wrapperPanic()  {
	defer func() {
		fmt.Println("a panic execute")
	}()
	bPanic()
}

func main() {
	wrapperPanic()
}
