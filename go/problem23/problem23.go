package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/divisors"
)

var (
	dc        divisors.Calc
	abundants []int
)

func main() {
	fmt.Println("Problem 23...")
	upperBound := 28123
	dc = divisors.NewCalc(upperBound)
	abundants = abundantNumsLessThan(upperBound)
	sum := 0
	for i := 1; i < upperBound; i++ {
		if !isSumOfAbundantNums(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}

func isSumOfAbundantNums(num int) bool {
	for i := 0; abundants[i] < num; i++ {
		for j := 0; abundants[i]+abundants[j] <= num; j++ {
			if abundants[i]+abundants[j] == num {
				return true
			}
		}
	}
	return false
}

func abundantNumsLessThan(num int) []int {
	nums := []int{}
	for i := 1; i < num; i++ {
		if isAbundant(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func isAbundant(num int) bool {
	return sumProperDivisors(num) > num
}

func sumProperDivisors(num int) int {
	sum := -num
	divisors, _ := dc.Divisors(num)
	for _, div := range divisors {
		sum += div
	}
	return sum
}
