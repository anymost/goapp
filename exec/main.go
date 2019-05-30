package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func lookPath()  {
	path, err := exec.LookPath("go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
}

func cmd()  {
	cmd := exec.Command("cat", "/Users/bradyzhang/Demo/goapp/micro/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func contextCmd() {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Nanosecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sleep", "5")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func start() {
	cmd := exec.Command("sleep", "10")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func main()  {
	// lookPath()
	// cmd()
	// contextCmd()
	start()
}
