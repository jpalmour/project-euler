package divisors

import (
	"github.com/jpalmour/project-euler/go/primes"
)

// TODO: speed up by using getPrimeFactorsMap
func GetDivisors(num int) []int {
	var divisors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func GetDivisorCount(num int) int {
	primes := primes.GetPrimeFactorsMap(num)
	divisorCount := 1
	for _, count := range primes {
		divisorCount *= count + 1
	}
	return divisorCount
}
