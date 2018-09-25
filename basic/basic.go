package main

import "fmt"

/**
类型
bool string int float
struct map array slice channel func

基本结构：
包名
导入
常量，变量，类型的定义
init函数
main函数
其他函数
 */



const age = 11
var name = "hello"
type IntType int

type Person struct {
	name string
	age int
}

func init()  {
	name = "jack"
}


func sayName()  {
	fmt.Println(name)
}

func sayAge()  {
	fmt.Println(age)
}

func (person Person) String() string  {
	return person.name
}


