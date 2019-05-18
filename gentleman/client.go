package main

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"log"
)

func main() {
	client := gentleman.New()
	client.URL("http://localhost:5000")
	req := client.Request()
	req.Path("/json")
	req.AddHeader("client", "go")
	res, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.String())
}
