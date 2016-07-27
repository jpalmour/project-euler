package primes

import (
	"fmt"
	"sort"
)

type Calc struct {
	UpperBound int
	Primes     []int
}

func NewCalc(upperBound int) Calc {
	primes := sieve(upperBound)
	return Calc{UpperBound: upperBound, Primes: primes}
}

func sieve(upperBound int) []int {
	sieve := make([]bool, upperBound)
	sieve[1] = true
	var primes []int
	for i := 2; i < upperBound; i++ {
		if sieve[i] == false {
			primes = append(primes, i)
			for j := 2 * i; j < upperBound; j += i {
				sieve[j] = true
			}
		}
	}
	return primes
}

func (c *Calc) IsPrime(num int) (bool, error) {
	if err := c.checkBounds(num); err != nil {
		return false, err
	}
	i := sort.SearchInts(c.Primes, num)
	return c.Primes[i] == num, nil
}

func (c *Calc) GetPrimeFactorsMap(num int) (map[int]int, error) {
	if err := c.checkBounds(num); err != nil {
		return nil, err
	}
	primeFactors := make(map[int]int)
	for _, p := range c.Primes {
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

func (c *Calc) checkBounds(num int) error {
	if num > c.UpperBound {
		return fmt.Errorf("%d too large for primes size %d", num, c.UpperBound)
	}
	return nil
}
