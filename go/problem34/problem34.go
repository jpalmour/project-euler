package main

import (
	"fmt"
)

var (
	digitFactorials = map[int]int{
		0: 1,
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
		6: 720,
		7: 5040,
		8: 40320,
		9: 362880,
	}
)

func main() {
	fmt.Println("Problem 34...")
	sum := 0
	for num := 3; num < 1000000000; num++ {
		if num == sumFactDigits(num) {
			sum += num
		}
	}
	fmt.Println(sum)
}

func sumFactDigits(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		num = num / 10
		sum += digitFactorials[digit]
	}
	return sum
}
