package main

import (
	"fmt"
	"math/big"
)

func main() {
	maxSum := 0
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			num := new(big.Int)
			sum := digitSum(num.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil))
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	fmt.Println(maxSum)
}

func digitSum(num *big.Int) int {
	sum := 0
	for _, digit := range num.String() {
		sum += int(digit) - '0'
	}
	return sum
}
