package main

import (
	"fmt"
	"strconv"
)

func handleSend(ch chan<- int)  {
	defer close(ch)
	for i := 0; i < 10; i++ {
		fmt.Println("send " + strconv.Itoa(i))
		ch <- i
		if i == 5 {
			break
		}
	}
}

func handleReceive(ch <-chan int)  {
	for {
		fmt.Println(<-ch)
	}
}


func main() {
	ch := make(chan int)
	go handleReceive(ch)
	handleSend(ch)
}
