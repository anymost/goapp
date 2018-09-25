package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

/**
1、 -arg=val 形式的参数用于flag.Bool这种方式去解析
2、 arg=val 的形式的参数用于flag.NArg()这种方式去解析
 */

func testArgs() {
	fmt.Println(os.Args[1:])
}

var Reload = flag.Bool("reload", false, "reload flag")
var Name = flag.String("name", "jack", "name flag")
var Time = flag.Int("time", 1, "time flag")

func testFlags() {
	flag.Parse()
	fmt.Println(*Reload)
	fmt.Println(*Name)
	fmt.Println(*Time)
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Arg(i))
	}
}

func readFileFromArgs()  {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		content, err := ioutil.ReadFile(flag.Arg(i))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(content))
	}
}



//func main() {
//	// testArgs()
//	// testFlags()
//	readFileFromArgs()
//}
