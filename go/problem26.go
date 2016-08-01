package main

import (
	"fmt"
)

func main() {
	fmt.Println("Problem 26...")
	maxCycleLength := 0
	maxNum := 100
	for num := 101; num < 1000; num++ {
		cycleLength := getCycleLength(num)
		if cycleLength > maxCycleLength {
			maxCycleLength = cycleLength
			maxNum = num
		}
	}
	fmt.Println(maxNum)
}

func getCycleLength(num int) int {
	remainders := map[int]bool{}
	remainder := 1
	for {
		remainder = (remainder % num) * 10
		if _, repeating := remainders[remainder]; repeating || remainder == 0 {
			break
		}
		remainders[remainder] = true
	}
	return len(remainders)
}
