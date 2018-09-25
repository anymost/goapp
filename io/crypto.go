package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hasher := sha256.New()
	fmt.Println(hasher.Sum([]byte("hello world")))
	fmt.Printf("%d\n", hasher.Sum([]byte("hello world")))
}
