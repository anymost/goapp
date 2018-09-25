package main

import (
	"net/http"
	"fmt"
)

func handleHttpRequest(response http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(response, "<h2>%s</h2>", request.URL.Path)
	response.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", handleHttpRequest)
	http.ListenAndServe("localhost:8776", nil)
}
