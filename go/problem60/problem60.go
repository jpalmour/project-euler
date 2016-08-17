package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
	"math"
	"strconv"
)

var pc primes.Calc

func PrimeConcatCombinations(set []int, choose int) [][]int {
	if choose == 0 {
		return [][]int{{}}
	}
	if len(set) < choose {
		return [][]int{}
	}
	combos := appendToFront(set[0], PrimeConcatCombinationsHelper(set[1:], choose-1, []int{set[0]}))
	combos = append(combos, PrimeConcatCombinations(set[1:], choose)...)
	return combos
}

func PrimeConcatCombinationsHelper(set []int, choose int, combMembers []int) [][]int {
	if choose == 0 {
		return [][]int{{}}
	}
	if len(set) < choose {
		return [][]int{}
	}
	isConcatablePrime := true
	for _, combMember := range combMembers {
		concat1, _ := strconv.Atoi(strconv.Itoa(combMember) + strconv.Itoa(set[0]))
		concat2, _ := strconv.Atoi(strconv.Itoa(set[0]) + strconv.Itoa(combMember))
		c1IsPrime, _ := pc.IsPrime(concat1)
		c2IsPrime, _ := pc.IsPrime(concat2)
		if !c1IsPrime || !c2IsPrime {
			isConcatablePrime = false
		}
	}
	combos := PrimeConcatCombinationsHelper(set[1:], choose, combMembers)
	if isConcatablePrime {
		combMembers = append(combMembers, set[0])
		combos = append(combos, appendToFront(set[0], PrimeConcatCombinationsHelper(set[1:], choose-1, combMembers))...)
	}
	return combos
}

func appendToFront(head int, nums [][]int) [][]int {
	res := [][]int{}
	for _, list := range nums {
		res = append(res, append([]int{head}, list...))
	}
	return res
}

func main() {
	pc = primes.NewCalc(1000000000)
	fmt.Println("done calculating primes")
	minSum := math.MaxInt32
	for _, combos := range PrimeConcatCombinations(pc.Primes[0:1500], 5) {
		fmt.Println(combos)
		sum := 0
		for _, prime := range combos {
			sum += prime
		}
		if sum < minSum {
			minSum = sum
		}
	}
	fmt.Println(minSum)
}
