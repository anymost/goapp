package main

import (
	"time"
	"fmt"
)

func TimeAndTick()  {
	tick := time.Tick(5e8)
	boom := time.After(1e9)
	for {
		select {
			case <- tick:
				fmt.Println("tick trigger")
			case <- boom:
				fmt.Println("time trigger")
		}
	}
}



func main() {
	// TimeAndTick()

	timer := time.NewTimer(1E9)
	ticker := time.NewTicker(1e9)
	fmt.Println(<-timer.C)
	fmt.Println(<-ticker.C)
}
