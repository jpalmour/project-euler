package main

import (
	"fmt"
)

func main() {
	fmt.Println("Problem 28")
	// top right:    (2n+1)^2
	// top left:     (2n+1)^2 -  (2n-1) - 1
	// bottom left:  (2n+1)^2 - 2(2n-1) - 2
	// bottom right: (2n+1)^2 - 3(2n-1) - 3
	//
	// total:       16n^2 + 4n + 4
	sum := 1
	for n := 1; n < 501; n++ {
		sum += 16*n*n + 4*n + 4
	}
	fmt.Println(sum)
}
