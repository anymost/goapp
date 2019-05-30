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
	ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)
	cmd := exec.CommandContext(ctx,"cat", "/Users/bradyzhang/Demo/goapp/micro/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//go func() {
	//	cancel()
	//}()
	_ = cmd.Run()
}

func main()  {
	// lookPath()
	// cmd()
	contextCmd()
}
