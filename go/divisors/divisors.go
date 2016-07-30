package divisors

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
	"sort"
)

type Calc struct {
	UpperBound int
	primes     []int
}

func NewCalc(upperBound int) Calc {
	return Calc{UpperBound: upperBound, primes: primes.NewCalc(upperBound).Primes}
}

func (c *Calc) Divisors(num int) ([]int, error) {
	if err := c.checkBounds(num); err != nil {
		return nil, err
	}
	if num < 1 {
		return []int{}, nil
	}
	return c.divisorsFromPrimeDivisors(c.PrimeDivisorsMap(num)), nil
}

func (c *Calc) DivisorsBrute(num int) ([]int, error) {
	divisors := []int{}
	if err := c.checkBounds(num); err != nil {
		return nil, err
	}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors, nil
}

func (c *Calc) DivisorCount(num int) (int, error) {
	if err := c.checkBounds(num); err != nil {
		return 0, err
	}
	divisorCount := 1
	for _, count := range c.PrimeDivisorsMap(num) {
		divisorCount *= count + 1
	}
	return divisorCount, nil
}

func (c *Calc) PrimeDivisorsMap(num int) map[int]int {
	primeFactors := make(map[int]int)
	for _, p := range c.primes {
		if num <= 1 {
			break
		}
		for num%p == 0 {
			num /= p
			primeFactors[p]++
		}
	}
	return primeFactors
}

func (c *Calc) divisorsFromPrimeDivisors(primes map[int]int) []int {
	if len(primes) == 0 {
		return []int{1}
	}
	values := [][]int{}
	i := 0
	for prime, count := range primes {
		row := []int{1}
		for j := 1; j <= count; j++ {
			row = append(row, prime*row[j-1])
		}
		values = append(values, row)
		i++
	}
	working := values[0]
	for rowNum := 1; rowNum < len(values); rowNum++ {
		tmp := []int{}
		for i := 0; i < len(working); i++ {
			for j := 0; j < len(values[rowNum]); j++ {
				tmp = append(tmp, working[i]*values[rowNum][j])
			}
		}
		working = tmp
	}
	sort.Ints(working)
	return working
}

func (c *Calc) checkBounds(num int) error {
	if num > c.UpperBound {
		return fmt.Errorf("%d too large for divisors.Calc size %d", num, c.UpperBound)
	}
	return nil
}
