package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/divisors"
)

func main() {
	fmt.Println("Problem 47")
	dc := divisors.NewCalc(1000000)
	for num := 1; true; num++ {
		if len(dc.PrimeDivisorsMap(num)) > 3 &&
			len(dc.PrimeDivisorsMap(num+1)) > 3 &&
			len(dc.PrimeDivisorsMap(num+2)) > 3 &&
			len(dc.PrimeDivisorsMap(num+3)) > 3 {
			fmt.Println(num)
			break
		}

	}
}
