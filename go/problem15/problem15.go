package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Problem 15...")
	pathCount := new(big.Int)
	pathCount.Binomial(40, 20)
	fmt.Println(pathCount)
}
