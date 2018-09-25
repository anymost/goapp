package pg

import (
	"../basic"
	"fmt"
	"math/big"
)

func main() {
	value := big.NewInt(111111)
	fmt.Println(value)
	a := big.NewInt(2)
	b := big.NewInt(3)
	c := big.NewInt(0)
	fmt.Println(c.Sub(a, b))
	basic.TestPackage()
}
