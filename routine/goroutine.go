package main

import (
	"fmt"
	"runtime"
	"time"
)

func longWait(ch chan string)  {
	time.Sleep(3e9)
	fmt.Println("long wait start")
	ch <- "long wait end"
}

func mediumWait(ch chan string)  {
	time.Sleep(2e9)
	fmt.Println("medium wait start")
	ch <- "medium wait end"
}

func shortWait(ch chan string)  {
	time.Sleep(1e9)
	fmt.Println("short wait start")
	ch <- "short wait end"
}

func main() {
	runtime.GOMAXPROCS(4)
	ch := make(chan string, 3)
	go longWait(ch)
	go mediumWait(ch)
	go shortWait(ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}
