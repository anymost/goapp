package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)



func readFromConsole()  {
	reader := bufio.NewReader(os.Stdin)
	for {
		val, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if strings.Trim(val, "\n") == "end" {
			break
		}
		fmt.Println(val)
	}
}

func readFileByLine()  {
	file, err := os.Open("./car.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		val, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(val)
	}
}

func readFileContent()  {
	content, err := ioutil.ReadFile("./car.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}




func readFileWithBuffer() {
	file, err := os.Open("./input.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := make([]byte, 2048)
	reader := bufio.NewReader(file)
	reader.Read(buffer)
	fmt.Println(string(buffer))
}





func readFileFromZip() {
	var r *bufio.Reader
	file, err := os.Open("test.go.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)

	if err != nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(reader)
	}

	for {
		value, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println(value)
	}
}
//
//func main() {
//	//readFromConsole()
//	// readFileByLine()
//	// readFileContent()
//	// readFileWithBuffer()
//	readFileFromZip()
//}
