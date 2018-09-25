package main

import (
	"fmt"
	"os"
)

/**
常量只能为四种基本类型 bool string int float
变量声明会默认赋予零值
:= 用于声明局部变量 局部变量未使用会报错
init 每个文件中都可以有该函数
 */

const a, b, c = 1, 2, 3
var d, e, f = 4, 5, 6

const (
	g = 7
	h = 8
	i = 9
	j = "h"
)

const (
	k  ConstType = iota
	l
	m
	n
)

type ConstType int

func sayConstType(val ConstType)  {
	fmt.Println(val)
}

func printPlatform()  {
	fmt.Println(os.Getenv("GOOS"))
}

func detectValueAndReference()  {
	a := 1
	b := a
	fmt.Println(&a)
	fmt.Println(&b)
}
