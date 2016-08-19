package main

import (
	"fmt"
	"math/big"
)

func main() {
	count := 0
	for n := int64(1); n < 25; n++ {
		for b := int64(1); true; b++ {
			pow := new(big.Int)
			pow.Exp(big.NewInt(b), big.NewInt(n), nil)
			limit := new(big.Int)
			limit.Exp(big.NewInt(10), big.NewInt(n), nil)
			if pow.Cmp(limit) > 0 {
				break
			}
			if len(pow.String()) == int(n) {
				count++
			}
		}
	}
	fmt.Println(count)
}
