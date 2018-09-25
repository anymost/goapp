package main

import (
	"fmt"
	"time"
)

func sendChan(ch chan<- int)  {
	count := 0
	for count < 10 {
		ch <- count
		count++
	}
}

func receiveChan(ch <-chan int)  {
	for {
		fmt.Println(<-ch)
	}
}


func main() {
	ch := make(chan int, 1)
	go sendChan(ch)
	go receiveChan(ch)
	time.Sleep(1e9)
}
