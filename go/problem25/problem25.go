package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(1)
	var i int
	for i = 0; len(a.String()) < 1000; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	fmt.Println(i)
}
