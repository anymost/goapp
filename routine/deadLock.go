package main

import (
	"fmt"
)

func f1(ch chan int)  {
	fmt.Println(<-ch)
}

func f2(ch chan int)  {
	ch <- 1
}


func main() {
	ch := make(chan int)
	// 这个会为什么会死锁呢？
	// 非常简单 f1 会阻塞主线程，一直等待，但是这个时候f2无法执行 导致死锁
	// 同时channel为有缓存的形式，则可以向里面发送，不用接收者准备好，那就不会死锁
	f2(ch)
	go f1(ch)
}
