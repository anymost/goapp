package main

import (
	"fmt"
	"bufio"
	"os"
)

func basicInput() {
	var (
		inputString string
		inputNumber int
		inputBool   bool
		)
	fmt.Println("please value")
	fmt.Scanln(&inputString, &inputNumber, &inputBool)
	fmt.Println(inputString, inputNumber, inputBool)
}

func formatInput() {
	input := "56  5212  hi"
	format := "%f  %d  %s"
	var (
		f float64
		i int
		s string
	)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println(f, i, s)
}

func testBuffer() {
	inputReader := bufio.NewReader(os.Stdin)
	val, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println(val)
	}
}
//
//func main() {
//	// basicInput()
//	testBuffer()
//}
