package main

import (
	"time"
	"fmt"
	"strconv"
)

func channelOdd(ch chan<- int)  {
	for i := 0; i < 10; i++ {
		if i % 2 == 1 {
			ch <- i
			time.Sleep(1e8)
		}
	}
}

func channelEven(ch chan<- int)  {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			ch <- i
			time.Sleep(1e8)
		}
	}
}

func selectTwoChannel(oddCh <-chan int, evenCh <-chan int)  {
	for {
		select {
		case a := <- oddCh:
			fmt.Println("odd is " + strconv.Itoa(a))
			if a == 7 {
				fmt.Println("break is running")
			}
		case b := <- evenCh:
			fmt.Println("even is " + strconv.Itoa(b))
		}
	}
}




func main() {
	chOne := make(chan int)
	chTwo := make(chan int)
	go channelOdd(chOne)
	go channelEven(chTwo)
	go selectTwoChannel(chOne, chTwo)
	time.Sleep(1e9)
}
