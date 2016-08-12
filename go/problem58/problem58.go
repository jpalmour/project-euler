package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
)

func main() {
	primeCount := 0.0
	total := 1.0
	pc := primes.NewCalc(1000000000)
	for n := 1; true; n++ {
		side := 2*n + 1
		for i := 0; i < 4; i++ {
			if isP, _ := pc.IsPrime(side*side - i*(side-1)); isP {
				primeCount++
			}
		}
		total += 4
		if primeCount/total < 0.1 {
			fmt.Println(side)
			break
		}
	}
}
