package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			for k := 0; k < 10; k++ {
				fmt.Printf("parent %d, child %d \n", i, k)
			}
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}
