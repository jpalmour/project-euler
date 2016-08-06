package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
)

func main() {
	fmt.Println("Problem 46")
	bound := 100000
	pc := primes.NewCalc(bound)
	nums := map[int]bool{}
	for _, prime := range pc.Primes {
		for square := 1; square < 1000; square++ {
			nums[prime+2*square*square] = true
		}
	}
	for n := 9; n < bound; n += 2 {
		prime, _ := pc.IsPrime(n)
		if !prime && !nums[n] {
			fmt.Println(n)
			break
		}
	}
}
