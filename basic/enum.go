package main

import "fmt"

const (
	INIT State = iota
	SUCCESS
	FAILED
)

type State int

func (state State) String() string {
	switch state {
	case INIT:
		return "init"
	case SUCCESS:
		return "success"
	case FAILED:
		return "failed"
	default:
		return "init"
	}
}

func main()  {
	fmt.Println(INIT)
	fmt.Println(SUCCESS)
	fmt.Println(FAILED)
}
