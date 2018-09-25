package main

import (
	"net/http"
	"fmt"
	"bufio"
	"io"
)

func main() {
	urls := []string{"https://www.baidu.com", "https://www.google.com"}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			break
		}
		reader := bufio.NewReader(resp.Body)
		for {
			value, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Println(value)
		}
	}
}
