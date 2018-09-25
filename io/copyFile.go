package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile()  {
	readFile, err := os.Open("./write.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writerFile, err := os.OpenFile("./read.txt", os.O_WRONLY | os.O_CREATE, 0666)
	defer writerFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(writerFile, readFile)
}

//func main() {
//	copyFile()
//}
