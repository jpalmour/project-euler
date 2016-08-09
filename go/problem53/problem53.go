package main

import (
	"fmt"
	"math/big"
)

func main() {
	ncr := new(big.Int)
	count := 0
	for n := 1; n <= 100; n++ {
		for r := 0; r <= n; r++ {
			ncr.Binomial(int64(n), int64(r))
			if ncr.Cmp(big.NewInt(1000000)) > 0 {
				fmt.Println(n)
				count++
			}
		}
	}
	fmt.Println(count)
}
