package main

import (
	"fmt"
)

func main() {
	fmt.Println("Problem 30...")
	sum := 0
	for num := 10; num < 10000000; num++ {
		if num == digit5PowerSum(num) {
			sum += num
		}
	}
	fmt.Println(sum)
}

func digit5PowerSum(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		num = num / 10
		sum += digit * digit * digit * digit * digit
	}
	return sum
}
