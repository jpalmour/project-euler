package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/divisors"
)

var (
	dc divisors.Calc
)

func main() {
	fmt.Println("Problem 21...")
	dc = divisors.NewCalc(10000)
	sum := 0
	for i := 1; i < 10000; i++ {
		if isAmicable(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}

func isAmicable(num int) bool {
	d := sumProperDivisors(num)
	return d != num && sumProperDivisors(d) == num
}

func sumProperDivisors(num int) int {
	sum := -num
	divisors, _ := dc.Divisors(num)
	for _, div := range divisors {
		sum += div
	}
	return sum
}
