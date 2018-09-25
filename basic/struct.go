package main

import (
	"fmt"
	"strconv"
	"reflect"
	"encoding/json"
)

/**
1 方法时作用在接收者上的函数，类型和其方法构成类，类型的所有方法称为方法集
2 对于访问结构体的值，不需要区分变量是值还是指针，会自动转换；结构体值方法可以通过指针访问，结构体指针方法不能通过值访问
3 结构体在内存中按块存储，性能相较于引用会更有优势
4 类型和方法不需要在同一个文件中，但需要在同一个包中
 */

type Man struct {
	name string
	age int
}

func structInit()  {
	var man1 Man
	man1.name = "jack"
	man1.age = 22

	man2 := Man{"jack", 22}
	fmt.Println(man2)

	man3 := &Man{name: "jack", age: 22}
	fmt.Println(man3)

	man4 := new(Man)
	man4 = &Man{"jack", 22}
	fmt.Println(man4)
}

func structConvert() {
	type Woman Man
	man := Man{"jack", 23}
	woman := Woman{"lucy", 22}
	woman = Woman(man)
	fmt.Println(reflect.TypeOf(woman))

}

func structMaker(name string, age int) *Man {
	return &Man{name, age}
}

func structTagged()  {
	type Car struct {
		Name string `json:"name"`
		Tire string `json:"tire"`
	}
	benz := Car{"benz", "continental"}
	val, _ := json.Marshal(benz)
	fmt.Println(string(val))

}


func structInner()  {
	type Inner struct {
		name string
	}

	type Outer struct {
		Inner
		age int
	}

	outer := Outer{Inner{"hello"}, 1}
	fmt.Println(outer.name)
}

func (man Man) sayName()  {
	fmt.Println(man.name)
}

func (man *Man) sayAge()  {
	fmt.Println(man.age)
}

func (man Man) String()string  {
	return man.name + strconv.Itoa(man.age)
}

func main()  {
	// structInit()
	// structConvert()
	// fmt.Println(structMaker("jack", 12))
	// structTagged()
	// structInner()
	//man := Man{"jack", 11}
	//man.sayName()
	//man.sayAge()
	// fmt.Println(Man{"hello", 11})
}
