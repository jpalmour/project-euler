package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Problem 16...")
	num := new(big.Int)
	num.Exp(big.NewInt(2), big.NewInt(1000), nil)
	sum := 0
	for _, digitStr := range strings.Split(num.String(), "") {
		digit, _ := strconv.Atoi(digitStr)
		sum += digit
	}
	fmt.Println(sum)
}
