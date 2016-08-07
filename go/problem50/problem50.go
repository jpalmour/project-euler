package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
)

func main() {
	fmt.Println("Problem 50...")
	pc := primes.NewCalc(1000000)
Loop:
	for count := len(pc.Primes); count > 0; count-- {
		for iStart := 0; iStart+count < len(pc.Primes); iStart++ {
			sum := 0
			for offset := 0; offset < count; offset++ {
				sum += pc.Primes[iStart+offset]
			}
			if sum > 999999 {
				break
			}
			if p, _ := pc.IsPrime(sum); p {
				fmt.Print(sum)
				break Loop
			}
		}
	}
}
