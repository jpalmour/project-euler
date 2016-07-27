package divisors

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
)

type Divisors struct {
}

func GetDivisorCount(num int, pc primes.Calc) int {
	primeDivisors, _ := getPrimeDivisorsMap(num, pc)
	divisorCount := 1
	for _, count := range primeDivisors {
		divisorCount *= count + 1
	}
	return divisorCount
}

func getPrimeDivisorsMap(num int, pc primes.Calc) (map[int]int, error) {
	if num > pc.UpperBound {
		return nil, fmt.Errorf("%d too large for primes size %d", num, pc.UpperBound)
	}
	primeFactors := make(map[int]int)
	for _, p := range pc.Primes {
		for num%p == 0 {
			num /= p
			primeFactors[p]++
		}
		if num == 1 {
			break
		}
	}
	return primeFactors, nil
}
