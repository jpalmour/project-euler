package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
	"strconv"
)

var pc primes.Calc

func main() {
	fmt.Println("Problem 49")
	pc = primes.NewCalc(100000)
	// TODO: I didn't read question carefully and thought digits couldn't be repeated.
	// A different approach would be much simpler now that I know digits can't be repeated.
	for _, combo := range combinations([]int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9}, 4) {
		primePerms := []int{}
		for _, perm := range permutations(combo) {
			num := digitsToInt(perm)
			p, _ := pc.IsPrime(num)
			if p && num > 999 {
				primePerms = append(primePerms, num)
			}
		}
		if len(primePerms) > 2 {
			seq := getArithmeticSequence(primePerms)
			if len(seq) > 0 {
				fmt.Println(seq)
				return
			}
		}
	}
}

func getArithmeticSequence(nums []int) string {
	fmt.Println(nums)
	for iGap := 1; iGap <= len(nums); iGap++ {
		for iStart := 0; iStart+2*iGap < len(nums); iStart++ {
			num1 := nums[iStart]
			num2 := nums[iStart+iGap]
			diff := num2 - num1
			num3 := num2 + diff
			if contains(nums, num3) && num1 != num2 && num1 != 1487 && num2 != 4817 {
				return fmt.Sprintf("%v%v%v", num1, num2, num3)
			}
		}
	}
	return ""
}

func contains(nums []int, val int) bool {
	for _, num := range nums {
		if num == val {
			return true
		}
	}
	return false
}

func digitsToInt(digits []int) int {
	numString := ""
	for _, digit := range digits {
		numString += strconv.Itoa(digit)
	}
	num, _ := strconv.Atoi(numString)
	return num
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
