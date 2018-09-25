package main

import (
	"fmt"
	"runtime"
)

func calculateValue(ch chan int)  {
	sum := 0
	for count:= 0; count < 9000000000; count++{
		sum += count
	}
	ch <- sum
}


func main() {
	runtime.GOMAXPROCS(3)
	ch := make(chan int)
	go calculateValue(ch)
	fmt.Println(<- ch)
}