package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/combinatorics"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Problem 20...")
	num := combinatorics.Factorial(big.NewInt(100))
	fmt.Println(num)
	sum := 0
	for _, digitStr := range strings.Split(num.String(), "") {
		digit, _ := strconv.Atoi(digitStr)
		sum += digit
	}
	fmt.Println(sum)
}
