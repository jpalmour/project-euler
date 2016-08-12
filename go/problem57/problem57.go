package main

import (
	"fmt"
	"math/big"
)

var cache = map[int]*big.Rat{}

func main() {
	count := 0
	for n := 1; n <= 1000; n++ {
		exp := getRoot2Exp(n)
		if numHasMoreDigits(exp) {
			count++
		}
	}
	fmt.Println(count)
}

func getRoot2Exp(n int) *big.Rat {
	sum := new(big.Rat)
	sum.SetInt64(1)
	inv := new(big.Rat)
	return sum.Add(sum, inv.Inv(getRoot2ExpDenom(n)))
}

func getRoot2ExpDenom(n int) *big.Rat {
	if n == 1 {
		val := new(big.Rat)
		val.SetInt64(2)
		cache[n] = val
	} else {
		sum := new(big.Rat)
		sum.SetInt64(2)
		inv := new(big.Rat)
		cache[n] = sum.Add(sum, inv.Inv(cache[n-1]))
	}
	return cache[n]
}

func numHasMoreDigits(num *big.Rat) bool {
	return len(num.Num().String()) > len(num.Denom().String())
}
