package main

import (
	"fmt"
	"strconv"
)



func sendWithBuffer(ch chan int) {
	count := 100
	for count > 0 {
		ch <- count
		fmt.Println("send " + strconv.Itoa(count))
		count--
	}
	close(ch)
}

func receiveWithBuffer(ch chan int) {
	for {
		value, ok := <- ch
		if ok {
			fmt.Println("receive " + strconv.Itoa(value))
		} else {
			break
		}
	}
}



func main() {
	chBuffer := make(chan int, 10)
	go receiveWithBuffer(chBuffer)
	sendWithBuffer(chBuffer)
}
