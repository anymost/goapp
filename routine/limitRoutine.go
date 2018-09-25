package main

import (
	"time"
	"fmt"
)

func printFrequency(ch <- chan time.Time)  {
	for {
		<- ch
		fmt.Println("hello world")
	}
}

func main() {
	tick := time.Tick(1e9)
	printFrequency(tick)
}
