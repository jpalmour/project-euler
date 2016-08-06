package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Problem 44...")
Loop:
	for j := 1; true; j++ {
		pj := (j * (3*j - 1)) / 2
		for k := j - 1; k > 0; k-- {
			pk := (k * (3*k - 1)) / 2
			D := pj - pk
			S := pj + pk
			if isPentagonal(D) && isPentagonal(S) {
				fmt.Println(D)
				break Loop
			}
		}
	}
	// TODO: above solution only finds one such D, but does not verify this D is the smallest (though it is)
	// This is a failed attempt that tries to verify the smallest D
	// for _, D := range pnc.PentagNumbers {
	// 	for j := 0; j+1 < len(pnc.PentagNumbers) && D > pnc.PentagNumbers[j+1]-pnc.PentagNumbers[j]; j++ {
	// 		pj := pnc.PentagNumbers[j]
	// 		pk := pj + D
	// 		pkIsPentag, _ := pnc.IsPentagonal(pk)
	// 		sumIsPentag, _ := pnc.IsPentagonal(pj + pk)
	// 		if pkIsPentag && sumIsPentag {
	// 			fmt.Println(D)
	// 			return
	// 		}
	// 	}
	// }
}

func isPentagonal(num int) bool {
	n := (math.Sqrt(24*float64(num)+1) + 1) / 6.0
	return float64(int(n)) == n
}
