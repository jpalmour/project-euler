package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	for n := 2; n < 10000; n++ {
		root := math.Sqrt(float64(n))
		if root != float64(int(root)) && period(n)%2 == 1 {
			count++
		}
	}
	fmt.Println(count)
}

func period(r int) int {
	wpBegin := int(math.Sqrt(float64(r)))
	wpNext, dNext := next(r, wpBegin, 1)
	period := 1
	for wpNext != wpBegin || dNext != 1 {
		wpNext, dNext = next(r, wpNext, dNext)
		period++
	}
	return period
}

func next(r, wp, d int) (wpNext, dNext int) {
	dNext = (r - wp*wp) / d
	aN := (int(math.Sqrt(float64(r))) + wp) / dNext
	wpNext = dNext*aN - wp
	return
}
