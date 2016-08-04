package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Problem 41...")
	pc := primes.NewCalc(1000000000)
	maxPrime := 2
	for _, prime := range pc.Primes {
		if isPandigital(prime) {
			maxPrime = prime
		}
	}
	fmt.Println(maxPrime)
}

func isPandigital(num int) bool {
	numStr := strconv.Itoa(num)
	for digit := 1; digit <= len(numStr); digit++ {
		if !strings.Contains(numStr, strconv.Itoa(digit)) {
			return false
		}
	}
	return true
}
