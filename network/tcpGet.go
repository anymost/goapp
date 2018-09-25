package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	var (
		count = 0
		data = make([]byte, 2048)
	)
	conn, err := net.Dial("tcp", "180.97.33.107:443")
	if err != nil {
		fmt.Println("connect " + err.Error())
		return
	}
	io.WriteString(conn, "GET / \n")
	for {
		count, err = conn.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(count)
		fmt.Printf(string(data[0: count]))
	}
	conn.Close()
}
