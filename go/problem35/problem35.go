package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
	"strconv"
)

func main() {
	fmt.Println("Problem 35")
	pc := primes.NewCalc(999999)
	count := 0
	for _, prime := range pc.Primes {
		if isCircularPrime(prime, pc) {
			count++
		}
	}
	fmt.Printf("Num circular primes: %v\n", count)
}

func isCircularPrime(num int, pc primes.Calc) bool {
	for _, val := range getRotations(num) {
		p, _ := pc.IsPrime(val)
		if !p {
			return false
		}
	}
	return true
}

func getRotations(num int) []int {
	str := strconv.Itoa(num)
	rotations := []int{num}
	for i := 1; i < len(str); i++ {
		str = str[1:] + str[0:1]
		num, _ = strconv.Atoi(str)
		rotations = append(rotations, num)

	}
	return rotations
}
