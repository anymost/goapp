package main

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleRequest(conn)
	}

}
func handleRequest(conn net.Conn)  {
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("connect is close")
			return
		}
		fmt.Println(string(buf))
	}
}
