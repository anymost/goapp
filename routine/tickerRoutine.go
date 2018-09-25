package main

import (
	"time"
	"fmt"
	"strconv"
)

func sendChannleEven(ch chan<- int)  {
	start := 0
	for {
		ch <- start
		start += 2
		time.Sleep(4e8)
	}
}

func sendChannelOdd(ch chan<- int)  {
	start := 1
	for {
		ch <- start
		start += 2
		time.Sleep(4e8)
	}
}

func selectDiffRoutine(chOne <- chan int, chTwo <- chan int, ticker *time.Ticker) {
	count := 0
	for {
		select {
		case a := <- chOne:
			count++
			fmt.Println("one " + strconv.Itoa(a))
		case b := <- chTwo:
			count++
			fmt.Println("two " + strconv.Itoa(b))
		case c := <- ticker.C:
			fmt.Println("ticker " + c.String())
		}
		if count == 20 {
			ticker.Stop()
		}
	}
}


func main() {
	chOne := make(chan int)
	chTwo := make(chan int)
	ticker := time.NewTicker(1e9)

	go sendChannleEven(chOne)
	go sendChannelOdd(chTwo)
	selectDiffRoutine(chOne, chTwo, ticker)
}
