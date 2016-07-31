package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
)

func main() {
	fmt.Println("Problem 27")
	bestA := -999
	bestB := -1000
	pc := primes.NewCalc(99999999)
	longestPrimeRun := 0
	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			run := getLongestRun(a, b, pc)
			if run > longestPrimeRun {
				bestA = a
				bestB = b
				longestPrimeRun = run
			}
		}
	}
	fmt.Println(bestA * bestB)
}

func quadratic(a, b, n int) int {
	return n*n + a*n + b
}

func getLongestRun(a, b int, pc primes.Calc) int {
	n := 0
	for prime := true; prime; n++ {
		val := quadratic(a, b, n)
		prime, _ = pc.IsPrime(val)
	}
	return n
}
