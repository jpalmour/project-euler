package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("Problem 13...")
	stringNums := strings.Split(input, "\n")
	sum := new(big.Int)
	for _, stringNum := range stringNums {
		i := new(big.Int)
		i, _ = i.SetString(stringNum, 10)
		sum.Add(sum, i)
	}
	fmt.Println(sum.String()[:10])
}
