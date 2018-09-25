package main

import "fmt"

func sendFor() chan int {
	ch := make(chan int)
	go func() {
		count := 0
		for count < 10{
			ch <- count
			count++
		}
		close(ch)
	}()
	return ch
}

func receiveFor(ch chan int)  {
	for value := range ch {
		fmt.Println(value)
	}
}

func main() {
	receiveFor(sendFor())
}
