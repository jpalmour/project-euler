package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Problem 32...")
	products := map[int]bool{}
	for lengthA := 1; lengthA < 8; lengthA++ {
		for lengthB := lengthA; lengthA+lengthB < 9; lengthB++ {
			getPandigitalProducts(lengthA, lengthB, products)
		}
	}
	sum := 0
	for product := range products {
		sum += product
	}
	fmt.Println(sum)
}

func getPandigitalProducts(lengthA, lengthB int, products map[int]bool) {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, combinationA := range combinations(digits, lengthA) {
		diff := difference(digits, combinationA)
		for _, combinationB := range combinations(diff, lengthB) {
			leftovers := difference(diff, combinationB)
			for _, a := range intPermutations(combinationA) {
				for _, b := range intPermutations(combinationB) {
					if digitMatch(a*b, leftovers) {
						products[a*b] = true
					}
				}
			}
		}
	}
}

func combinations(nums []int, length int) [][]int {
	if length == 0 {
		return [][]int{{}}
	}
	if len(nums) == length {
		return [][]int{nums}
	}
	combos := appendToFront(nums[0], combinations(nums[1:], length-1))
	combos = append(combos, combinations(nums[1:], length)...)
	return combos
}

func appendToFront(head int, nums [][]int) [][]int {
	res := [][]int{}
	for _, list := range nums {
		res = append(res, append([]int{head}, list...))
	}
	return res
}

func difference(numsA, numsB []int) []int {
	diff := []int{}
	for _, numA := range numsA {
		found := false
		for _, numB := range numsB {
			if numB == numA {
				found = true
			}
		}
		if !found {
			diff = append(diff, numA)
		}
	}
	return diff
}

func intPermutations(digits []int) []int {
	perms := []int{}
	for _, perm := range permutations(digits) {
		perms = append(perms, digitsToInt(perm))
	}
	return perms
}

func permutations(digits []int) [][]int {
	if len(digits) == 0 {
		return [][]int{}
	}
	if len(digits) == 1 {
		return [][]int{digits}
	}
	perms := [][]int{}
	for i, digit := range digits {
		dup := make([]int, len(digits))
		copy(dup, digits)
		withoutDigit := append(dup[0:i], dup[i+1:]...)
		permsWithoutDigit := permutations(withoutDigit)
		perms = append(perms, appendToFront(digit, permsWithoutDigit)...)
	}
	return perms
}

func digitsToInt(digits []int) int {
	numString := ""
	for _, digit := range digits {
		numString += strconv.Itoa(digit)
	}
	num, _ := strconv.Atoi(numString)
	fmt.Println(num)
	return num
}

func digitMatch(num int, digits []int) bool {
	numDigits := []int{}
	for num != 0 {
		numDigits = append(numDigits, num%10)
		num /= 10
	}
	sort.Ints(numDigits)
	sort.Ints(digits)
	return reflect.DeepEqual(numDigits, digits)
}
