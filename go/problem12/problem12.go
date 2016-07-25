package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/divisors"
	"github.com/jpalmour/project-euler/go/triangleNumbers"
)

func main() {
	fmt.Println("Problem 12...")
	maxDivisors := 0
	for t := triangleNumbers.New(); true; t.Next() {
		numDivisors := divisors.GetDivisorCount(t.Value)
		if numDivisors > maxDivisors {
			maxDivisors = numDivisors
			fmt.Printf("%v has %v divisors\n", t.Value, numDivisors)
		}
		if numDivisors > 500 {
			fmt.Println(t.Value)
			return
		}
	}
}
