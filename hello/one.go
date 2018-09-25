// 变量
package hello

import (
	"fmt"
)

//var a, b int
//var a int

func test1() {
	var a, b int
	a = 1
	b = 1
	fmt.Println(a)
	_ = b
}

func test2() {
	var a, b = 1, 2
	fmt.Println(a, b)

}

func test3() {
	a, b := 1, 2
	fmt.Println(a, b)
}

func test4() {
	var (
		a int
		b string
	)
	a = 1
	b = "hi"
	fmt.Println(a, b)
}

var helloType = "hello type"
