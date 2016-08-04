package main

import (
	"fmt"
)

func main() {
	fmt.Println("problem39...")
	maxCount := 0
	maxP := 3
	for p := 3; p < 1001; p++ {
		count := 0
		for a := 1; a < p; a++ {
			for b := 1; b < p-a; b++ {
				if a*a+b*b == (p-a-b)*(p-a-b) {
					count++
				}
			}
		}
		if count > maxCount {
			maxCount = count
			maxP = p
		}
	}
	fmt.Println(maxP)
}
