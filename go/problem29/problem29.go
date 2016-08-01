package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/divisors"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Problem 29...")
	dc := divisors.NewCalc(1000)
	distinctTerms := map[string]bool{}
	for a := 2; a < 101; a++ {
		for b := 2; b < 101; b++ {
			aFactored := dc.PrimeDivisorsMap(a)
			for prime, count := range aFactored {
				aFactored[prime] = count * b
			}
			distinctTerms[toString(aFactored)] = true
		}
	}
	fmt.Println(len(distinctTerms))
}

func toString(primes map[int]int) string {
	var keys []int
	for k := range primes {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var list []string
	for _, k := range keys {
		list = append(list, fmt.Sprintf("prime: %v, count: %v", k, primes[k]))
	}
	str := strings.Join(list, "; ")
	return str
}
