package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/divisors"
	"github.com/jpalmour/project-euler/go/triangleNumbers"
)

func main() {
	fmt.Println("Problem 12...")
	dc := divisors.NewCalc(100000000)
	maxDivisors := 0
	for t := triangleNumbers.New(); true; t.Next() {
		numDivisors, _ := dc.DivisorCount(t.Value)
		if numDivisors > maxDivisors {
			maxDivisors = numDivisors
		}
		if numDivisors > 500 {
			fmt.Println(t.Value)
			return
		}
	}
}
