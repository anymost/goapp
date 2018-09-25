package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
引用类型默认是按引用传递函数参数
1 basic
2 命名返回值
3 使用...展开/收集参数
4 defer 多个defer逆序执行 用于关闭文件 关闭数据库连接  解锁加锁的资源 打印最后的记录
5 递归
6 函数作为参数
7 函数作为返回值
8 匿名执行函数/闭包
   (1) 匿名函数配合defer，用于修改函数的命名返回值，可以在return之后再次修改返回值
   (2) runtime.caller(1) 用于定位在执行的文件及行数
   (3) 使用内存缓存 把计算过的值缓存起来
*/
/**
  1 basic
  2 命名返回值
  3 ...收集参数
  4 defer
  5 递归
  6 高阶函数
  7 匿名函数
 */

type hello func(int, int) int

var result [42]int


func baseFunc(a int) int {
	return a
}

func returnNamedValueFunc() (a int) {
	a = 1
	return
}

func pointerFunc(a, b int, c *int) {
	*c = a + b
}

func floatArgs(a ...int) {
	fmt.Println(len(a))
}


func deferFuc() {
	a := 0
	defer fmt.Println(a)
	a++
	return

}

func fibonacci(start int)(result int)  {
	if start <= 1 {
		result = 1
	} else {
		result = fibonacci(start - 1) + fibonacci(start - 2)
	}
	return
}

func betterFibonacci(start int)(val int) {
	if result[start] != 0 {
		val = result[start]
		return
	}
	if start <= 1 {
		val = 1
	} else {
		val = betterFibonacci(start-1) + betterFibonacci(start-2)
	}
	result[start] = val
	return
}

func funAsArg(x int, y int, fn func(int, int) int)  {
	fmt.Println(fn(x, y))
}

var noNameFunc = func(x, y int) int {
	return x + y
}

func Instant() {
	func(x ...int) {
		fmt.Println(x)
	}(1, 2, 3, 4)
}

func returnFunc(a int) func(int) int {
	return func(i int) int {
		return a + i
	}
}

func bibao() func(int) int {
	x := 0
	return func(i int) int {
		x += i
		return x
	}
}

func where() {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(file, line)
}

func checkTime(num int, callback func(int)int)  {
	start := time.Now()
	callback(num)
	end := time.Now()
	fmt.Println(end.Sub(start))
}

type funcType func(a int, b int)int

func basicFunc(a, b int)int  {
	return a + b
}


func main() {
	checkTime(41, fibonacci)
	checkTime(41, betterFibonacci)
}


