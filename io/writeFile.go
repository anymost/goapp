package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFileFromConsole()  {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello world")
	}
	writer.Flush()
}

func writeFile()  {
	file, err := os.OpenFile("./write.txt", os.O_WRONLY | os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 1000; i++ {
		writer.WriteString("hello world\n")
	}
	writer.Flush()
}



//func main() {
//	// writeFileFromConsole()
//	writeFile()
//}
