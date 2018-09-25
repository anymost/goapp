package hi

import (
	"fmt"
	"strings"
	"time"
)

var printLn = fmt.Println

func main() {
	a, b := "hello", "world"
	_, _ = a, b
	printLn(strings.HasPrefix(a, "h"))
	array := strings.Fields("HELLO WORLD HELLO WORLD")
	printLn(strings.ToLower("HELLO WORLD"))
	for _, val := range array {
		printLn(val)
	}
	newString := strings.NewReader("hello world i am unhappy")
	printLn(newString)
	date := time.Now()
	fmt.Printf("%d-%d-%d", date.Year(), date.Month(), date.Day())

}
