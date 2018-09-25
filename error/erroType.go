package main

import (
	"errors"
	"fmt"
)

func customError()  {
	err := errors.New("bad emotion")
	fmt.Println(err)
}

func initPanic()  {
	fmt.Println("start")
	panic("crash")
	fmt.Println("end")
}


