package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Problem 43")
	sum := 0
	for _, perm := range permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		if hasDivProp(perm) {
			sum += digitsToInt(perm)
		}
	}
	fmt.Println(sum)
}

func hasDivProp(digits []int) bool {
	return len(digits) > 9 &&
		getNum(1, digits)%2 == 0 &&
		getNum(2, digits)%3 == 0 &&
		getNum(3, digits)%5 == 0 &&
		getNum(4, digits)%7 == 0 &&
		getNum(5, digits)%11 == 0 &&
		getNum(6, digits)%13 == 0 &&
		getNum(7, digits)%17 == 0
}

func getNum(i int, nums []int) int {
	a := strconv.Itoa(nums[i])
	b := strconv.Itoa(nums[i+1])
	c := strconv.Itoa(nums[i+2])
	num, _ := strconv.Atoi(a + b + c)
	return num
}

func digitsToInt(digits []int) int {
	numString := ""
	for _, digit := range digits {
		numString += strconv.Itoa(digit)
	}
	num, _ := strconv.Atoi(numString)
	return num
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

func appendToFront(head int, nums [][]int) [][]int {
	res := [][]int{}
	for _, list := range nums {
		res = append(res, append([]int{head}, list...))
	}
	return res
}
