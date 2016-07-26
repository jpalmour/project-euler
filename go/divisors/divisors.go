package divisors

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
)

type Divisors struct {
}

func GetDivisorCount(num int, l primes.List) int {
	primeDivisors, _ := getPrimeDivisorsMap(num, l)
	divisorCount := 1
	for _, count := range primeDivisors {
		divisorCount *= count + 1
	}
	return divisorCount
}

func getPrimeDivisorsMap(num int, l primes.List) (map[int]int, error) {
	if num > l.Bound {
		return nil, fmt.Errorf("%d too large for primes size %d", num, l.Bound)
	}
	primeFactors := make(map[int]int)
	for _, p := range l.Primes {
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
