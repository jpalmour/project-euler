package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Problem 48")
	sum := big.NewInt(0)
	for num := int64(1); num < 1001; num++ {
		var power big.Int
		power.Exp(big.NewInt(num), big.NewInt(num), nil)
		sum.Add(sum, &power)
	}
	numStr := sum.String()
	fmt.Println(numStr[len(numStr)-10:])
}
