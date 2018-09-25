package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
	1 map key可以是string,int,float; value 可以是any，
      可以使用len来获取map的key的数量，也可以在make map中指定初始容量
	2 map 是引用类型，赋值传递的是引用
    3 make(map[int]string)等价于map[int]string{}
    4 对于map， array，必须进行初始化，对于struct 无需进行初始化
*/

func basicMap() {
	// map1 := make(map[int]string)
	map2 := map[int]string{1: "hello", 2: "world"}
	map3 := map2
	map2[1] = "jack"
	fmt.Println(map2)
	fmt.Println(map3)
}

func hasKeyMap() {
	map3 := map[string]string{"first": "first", "second": "second"}
	val, isExist := map3["first"]
	if isExist {
		fmt.Println(val)
	}
}

func deleteKeyMap() {
	map4 := make(map[int]int, 10)
	map4[1] = 1
	map4[2] = 2
	fmt.Println(map4)
	delete(map4, 1)
	fmt.Println(map4)
}

func multiMap() {
	map5 := make(map[int]map[string]string)
	map5[1] = map[string]string{"first": "first"}
	map5[2] = map[string]string{"second": "second"}
	fmt.Println(map5)
}

func sortMap() {
	map6 := map[string]int{
		"h": 6,
		"e": 5,
		"a": 4,
		"b": 3,
		"k": 2,
		"x": 1,
		"z": 0,
	}
	keyList := make([]string, len(map6))
	count := 0
	for key := range map6 {
		keyList[count] = key
		count++
	}
	sort.Strings(keyList)
	for _, key := range keyList {
		fmt.Println("key is " + key + " value is " + strconv.Itoa(map6[key]))
	}
}

func sortMapFunc() {
	personMap := map[string]string{
		"lucy": "lucy woman",
		"jack": "jack man",
	}
	personSlice := make([]string, len(personMap))
	index := 0
	for key := range personMap {
		personSlice[index] = key
		index++
	}
	sort.Strings(personSlice)
	for _, name := range personSlice {
		fmt.Println("name is " + name + " value is " + personMap[name])
	}
}

func main() {
	// basicMap()
	// initMap()
	// hasKeyMap()
	sortMapFunc()
	// deleteKeyMap()
	// multiMap()
	// sortMap()
	// sortMapFunc()
}
