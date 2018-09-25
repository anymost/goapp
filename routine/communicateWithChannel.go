package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "hello"
	ch <- "world"
}

func receiveData(ch chan string)  {
	var input string
	for {
		input = <- ch
		fmt.Println(input)
	}
}

func main() {
	ch := make(chan string)
	go receiveData(ch)
	sendData(ch)
	// time.Sleep(1e9)
}
