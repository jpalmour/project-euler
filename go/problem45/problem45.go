package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Problem 45")
	for i := 144; true; i++ {
		hex := i * (2*i - 1)
		if isPentagonal(hex) {
			fmt.Println(hex)
			break
		}
	}
}

func isPentagonal(num int) bool {
	n := (math.Sqrt(24*float64(num)+1) + 1) / 6.0
	return float64(int(n)) == n
}
