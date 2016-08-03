package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
)

func main() {
	fmt.Println("Problem 37...")
	pc := primes.NewCalc(100000000)
	sum := -17
	for _, prime := range pc.Primes {
		if isTruncatable(prime, pc) {
			sum += prime
		}
	}
	fmt.Println(sum)
}

func isTruncatable(num int, pc primes.Calc) bool {
	return isRightTruncatable(num, pc) && isLeftTruncatable(num, pc)
}

func isRightTruncatable(num int, pc primes.Calc) bool {
	prime, _ := pc.IsPrime(num)
	next := num / 10
	return prime && (num < 10 || isRightTruncatable(next, pc))
}

func isLeftTruncatable(num int, pc primes.Calc) bool {
	prime, _ := pc.IsPrime(num)
	mod := 1
	dup := num
	for dup > 10 {
		dup /= 10
		mod *= 10
	}
	next := num % mod
	return prime && (num < 10 || isLeftTruncatable(next, pc))
}
