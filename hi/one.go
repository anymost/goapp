package hi

import (
	//"../hello"
	"fmt"
	"math"
	"math/rand"
	//"os"
)

var print = fmt.Println

//var printAll = fmt.Println
//
//func test(name string, age int) (helloName string, helloAge int) {
//	return name + string(age), age
//}
//
//func returnInt() int {
//	return 1
//}
//
//func helloTest(name string, age int)(returnName string, returnInt int) {
//	return name, age
//}
//
//func result() (value int) {
//	const name = "hello world"
//	fmt.Println(name)
//	age := 1
//	_ = age
//	return 1
//}
//
//func init() {
//	fmt.Println("init package")
//	fmt.Println(returnInt())
//}

//func hi() {
//const (
//	a = iota
//	b
//	c
//	d
//	e
//)
//
//fmt.Println(result())
//hello.SayHello()
//helloTest("hello", 1)
//printAll("hello world")
//fmt.Println(a)
//
//a, b := 1, "hello"
//fmt.Println(a)
//fmt.Println(b)
//var c int
//c = 1
//fmt.Println(c)
//	fmt.Println(os.Getuid())
//	fmt.Println(os.Getenv("GOOS"))
//	fmt.Printf("%s hhaha %d", "hello", 1)
//	test := 1
//	fmt.Println(test)
//}
//
//var print = fmt.Println
//var a = "G"
//
//func hi() {
//	n()
//	m()
//	n()
//}
//
//func n() { print(a) }
//
//func m() {
//	a := "O"
//	print(a)
//}
//
//var a string
//
//func hi() {
//	a = "G"
//	print(a)
//	f1()
//}
//
//func f1() {
//	a := "O"
//	print(a)
//	f2()
//}
//
//func f2() {
//	print(a)
//}
func main() {
	a, b := 1, 1
	var c = 2.3
	print(a == b)
	print(c)
	print(math.Floor(c))

	//var d int32
	//var e = 5
	//d += 5
	//d := 1
	//const e = 3

	name := 1 + 1i
	print(name)

	print(math.MaxInt8)

	type ByteSize float64
	const (
		_           = iota // 通过赋值给空白标识符来忽略值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(TB)

	print(rand.Intn(100))

	for i := 0; i < 1000; i++ {
		print(i)
	}

	type Rope string
	var role Rope = "hello world"
	print(role)

	var test Rope = "hi world"
	print(test + role)

	print(len(test))
}
