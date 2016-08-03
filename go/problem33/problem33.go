package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Problem 33...")
	product := new(big.Rat)
	product.SetInt64(1)
	for i := 10; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			it := i / 10
			iu := i % 10
			jt := j / 10
			ju := j % 10
			num := big.NewRat(int64(i), int64(j))
			if it == ju && jt != 0 {
				if num.Cmp(big.NewRat(int64(iu), int64(jt))) == 0 {
					product.Mul(product, num)
				}
			}
			if iu == jt && ju != 0 {
				if num.Cmp(big.NewRat(int64(it), int64(ju))) == 0 {
					product.Mul(product, num)
				}
			}
		}
	}
	fmt.Println(product.Denom())
}
