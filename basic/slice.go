package main

import (
	"fmt"
	"sort"
	"reflect"
)

/**
1 数组是值类型，长度不可变
2 数组的声明创建方式 [10]int、 [...]int
 */

 /**
 1 切片是引用类型，长度可变，通过指向相关数组的指针,len,cap来确定
 2 切片的声明创建方式 []int [10]int[:]
  */

  /**
  	make([]int, 50, 100) make返回类型T的初始值
  	new([100]int)[0:50] new返回一个指针，指向T的零值
    array := append(array, 1, 2)
    copy(array1, array2)
   */



func arrayInit()  {
	var array [10]int
	fmt.Println(array)
}

func arrayValue()  {
	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array2 := array1
	array2[0] = 2
	fmt.Println(array1)
}

func forRangeArray() {
	array := [4]int{1, 2, 3, 4}
	for index, value := range array {
		fmt.Println(index)
		fmt.Println(value)
	}
}

func initArray() {
	array1 := [10]int{1, 2, 3}
	array2 := [...]int{1, 2, 3}
	array3 := []int{1, 2, 3}
	fmt.Print(array1, array2, array3)
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(array3))
}

func matrix()  {
	var array [10][10]int
	for i := 0; i < 10; i++ {
		for k := 0; k < 10; k++ {
			array[i][k] = i + k
		}
	}
	fmt.Println(array)
}

func initialArray(arr [10]int) int  {
	result := 0
	for _, val := range arr {
		result += val
	}
	return result
}

func initialSlice(arr []int) int  {
	result := 0
	for _, val := range arr {
		result += val
	}
	return result
}

func arrayCap()  {
	array := [5]int{1, 2, 3, 4, 5}
	newArray := array[2: 3]
	fmt.Println(cap(newArray))
	fmt.Println(newArray)
}

func reslice()  {
	arr := make([]int, 5, 10)
	fmt.Println(arr)
	arr = arr[:10]
	fmt.Println(arr)
}



func copyOrAppend()  {
	slice1 := []int{1, 3, 4}
	slice2 := make([]int, 10)
	copy(slice2, slice1)
	fmt.Println(slice2)
	slice2 = append(slice2, 3, 4, 5)
	fmt.Println(slice2)
}

func appendSlice()  {
	array := []int{1, 2, 3}
	array2 := []int{4, 5, 6}
	for _, val := range array2 {
		array = append(array, val)
	}
	fmt.Println(array)
}

//func sliceApply() {
//	str := "hello"
//	for _, value := range str {
//		fmt.Println(value)
//	}
//	strArray := []byte(str)
//	strArray[0] = 'c'
//	str = string(strArray)
//	fmt.Println(str)
//}

//func makeAndNew()  {
//	array1 := make([]int, 10)
//	array1 = []int{1, 2, 3, 4}
//	fmt.Println(array1)
//	array2 := new([10]int)[:]
//	array2 = []int{1, 2, 3, 4}
//	fmt.Println(array2)
//}

func stringSlice()  {
	str := "好hello world"
	fmt.Println([]byte(str))
}

func modifyString()  {
	str := "hello"
	fmt.Println(str[0:1]+str[1:2])
}

func sortAndSearchArray()  {
	arr := []int{7, 6, 5, 4, 3, 2, 1}
	sort.Ints(arr)
	fmt.Println(sort.IntsAreSorted(arr))
	// 需要先排序，再查找
	fmt.Println(sort.SearchInts(arr, 4))
}

func main()  {
	//array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//start := time.Now()
	//initialArray(array)
	//middle := time.Now()
	//fmt.Println(middle.Sub(start))
	//middle = time.Now()
	//initialSlice(slice)
	//end := time.Now()
	//fmt.Println(end.Sub(middle))
	// modifyString()
	// sortAndSearchArray()
	appendSlice()
 }
