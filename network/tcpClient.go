package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	client, _ := net.Dial("tcp", ":50000")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input user name")
	userName, _ := inputReader.ReadString('\n')
	userName = strings.Trim(userName, "\n")
	for {
		fmt.Println("what do you want to say?")
		content, _ := inputReader.ReadString('\n')
		if content == "quit\n" {
			client.Close()
			return
		}
		client.Write([]byte(userName + " says " + content))
	}

}
