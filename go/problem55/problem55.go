package main

import (
	"fmt"
	"math/big"
)

func main() {
	// TODO: not sure if 0 counts
	count := 0
	for num := 10; num < 10000; num++ {
		if isLychrel(num) {
			count++
		}
	}
	fmt.Println(count)
}

func isLychrel(num int) bool {
	iterations := 1
	val := big.NewInt(int64(num))
	for iterations < 50 {
		val = next(val)
		if isPalindrome(val) {
			return false
		}
		iterations++
	}
	return true
}

func next(num *big.Int) *big.Int {
	revStr := ""
	for _, digit := range num.String() {
		revStr = string(digit) + revStr
	}
	ret := new(big.Int)
	ret.SetString(revStr, 10)
	return ret.Add(ret, num)
}

func isPalindrome(num *big.Int) bool {
	revStr := ""
	for _, digit := range num.String() {
		revStr = string(digit) + revStr
	}
	return num.String() == revStr
}
