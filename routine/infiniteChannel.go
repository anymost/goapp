package main

import "fmt"

func infiniteSend(ch chan int)  {
	count := 0
	for {
		ch <- count
		count++
	}
}

func infiniteReceive(ch chan int)  {
	for {
		fmt.Println(<- ch)
	}
}


func main() {
	ch := make(chan int)
	go infiniteSend(ch)
	infiniteReceive(ch)
}
