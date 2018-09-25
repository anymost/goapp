package main

import (
	"fmt"
	"reflect"
)


type Person struct {
	Name string
	Age int
	Address string
}

func (person Person) Generate(prefix string) {
	fmt.Println(prefix + person.Name)
}

func reflectTypeOrValue()  {
 	a := 1.1
 	fmt.Println(reflect.TypeOf(a))
 	valueA := reflect.ValueOf(a)
 	fmt.Println(valueA)
 	fmt.Println(valueA.Type())
 	fmt.Println(valueA.Kind())
 	fmt.Println(valueA.Interface())
 	fmt.Println(valueA.Float())
}


func setValueFromReflect()  {
	a := 1.2
	valueA := reflect.ValueOf(&a)
	valueA = valueA.Elem()
	if valueA.CanSet() {
		valueA.SetFloat(2.3)
	}
	fmt.Println(a)
}

func reflectStruct()  {
	val := Person{"jack", 23, "China"}
	valA := reflect.ValueOf(val)
	fmt.Println(valA.Kind())
	for i := 0; i < valA.NumField(); i++ {
		fmt.Println(valA.Field(i))
	}
	fmt.Println(valA.NumMethod())
	valA.Method(0).Call([]reflect.Value{reflect.ValueOf("HELLO")})
}

func modifyStructByReflect()  {
	val := Person{"jack", 11, "China"}
	valA := reflect.ValueOf(&val).Elem()
	valA.Field(0).SetString("ruby")
	fmt.Println(valA)
}



func reflectMethod()  {
	sayHello := func(prefix string)string  {
		return prefix + "jack"
	}
	helloMethod := reflect.ValueOf(sayHello)
	value := helloMethod.Call([]reflect.Value{reflect.ValueOf("hello")})
	fmt.Println(value[0])
}


func main() {
	// reflectTypeOrValue()
	// setValueFromReflect()
	// reflectStruct()
	// reflectMethod()
	modifyStructByReflect()
}
