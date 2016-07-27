package main

import (
	"fmt"
)

func main() {
	fmt.Println("Problem 14...")
	fmt.Println(getMaxChain(1000000))
}

func getMaxChain(num int) int {
	maxChainLength := 0
	maxChainInput := 1
	for i := 1; i < num; i++ {
		value := getChainSize(i)
		if value > maxChainLength {
			maxChainInput = i
			maxChainLength = value
		}
	}
	return maxChainInput
}

func getChainSize(num int) int {
	if num == 1 {
		return 1
	} else {
		return getChainSize(collatz(num)) + 1
	}
}

func collatz(num int) int {
	if num%2 == 0 {
		return num / 2
	}
	return 3*num + 1
}
