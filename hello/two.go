// 常量
package hello

import "fmt"

func test5() {
	const a = 1
	fmt.Println(a)

	const (
		b, c = 10, 20
		d    = false
	)
	fmt.Println(b, c, d)
}
