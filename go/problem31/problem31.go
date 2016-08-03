package main

import (
	"fmt"
)

func main() {
	fmt.Println("Problem 31...")
	count := 1
	for a := 0; a < 201; a += 100 {
		for b := 0; b < 201; b += 50 {
			for c := 0; c < 201; c += 20 {
				for d := 0; d < 201; d += 10 {
					for e := 0; e < 201; e += 5 {
						for f := 0; f < 201; f += 2 {
							for g := 0; g < 201; g += 1 {
								if a+b+c+d+e+f+g == 200 {
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}
