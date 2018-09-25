package main

/**
通过 val.(type) 可以检查接口变量是否是某一个具体类型
也可以检查变量是否实现了某个接口
 */

import (
	"fmt"
	"reflect"
)

type Singer interface {
	sing() string
}

type MaleSinger struct {
	name string
	song string
}

type FemaleSinger struct {
	name string
	song string
}

func main() {
	// typeDetect()
	//typeSwitchDetect()
	//values := []interface{}{1, 2.2, "hello", true}
	//multiTypeSwitch(values)
}

func typeSwitchDetect() {
	singers := []Singer{
		&MaleSinger{"jack", "my heart will go on"},
		&FemaleSinger{"lucy", "fireflies"},
	}
	for _, singer := range singers {
		switch t := singer.(type) {
		case *MaleSinger:
			fmt.Printf("type %T, value %v\n", t, t)
		case *FemaleSinger:
			fmt.Printf("type %T, value %v\n", t, t)
		default:
			fmt.Printf("type %T, value %v\n", t, t)
		}
	}
}

func multiTypeSwitch(values []interface{})  {
	for  _, value := range values {
		switch value.(type) {
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float64")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("string")
		default:
			fmt.Println("other type")
		}
	}
}

func typeDetect() {
	singers := []Singer{
		&MaleSinger{"jack", "my heart will go on"},
		&FemaleSinger{"lucy", "fireflies"},
	}
	for _, singer := range singers {
		fmt.Println(reflect.TypeOf(singer))
		if val, ok := singer.(*MaleSinger); ok {
			fmt.Println(reflect.TypeOf(val))
		}
		if val, ok := singer.(*FemaleSinger); ok {
			fmt.Println(reflect.TypeOf(val))
		}
	}
}

func (singer *MaleSinger) sing() string {
	return "male singer: " + singer.name + " " + singer.song
}

func (singer *FemaleSinger) sing() string {
	return "female singer: " + singer.name + " " + singer.song
}
