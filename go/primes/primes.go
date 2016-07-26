package primes

import (
	"fmt"
)

type List struct {
	Bound  int
	Primes []int
	sieve  []bool
}

func NewList(bound int) List {
	sieve := make([]bool, bound)
	sieve[1] = true
	var primes []int
	for i := 2; i < bound; i++ {
		if sieve[i] == false {
			primes = append(primes, i)
			for j := 2 * i; j < bound; j += i {
				sieve[j] = true
			}
		}
	}
	return List{sieve: sieve, Bound: bound, Primes: primes}
}

func (l *List) IsPrime(num int) (bool, error) {
	if num > l.Bound {
		return false, fmt.Errorf("%d too large for primes size %d", num, l.Bound)
	}
	return !l.sieve[num], nil
}

func (l *List) GetPrimeFactorsMap(num int) (map[int]int, error) {
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
