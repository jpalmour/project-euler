package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Problem 38...")
	max := int64(0)
	for n := 2; n < 9; n++ {
		for a := 1; a < 10000; a++ {
			concat := strconv.Itoa(a)
			for num := 2; num <= n; num++ {
				concat += strconv.Itoa(num * a)
			}
			if isPandigital(concat) {
				num, _ := strconv.ParseInt(concat, 10, 64)
				if num > max {
					max = num
				}
			}
		}
	}
	fmt.Println(max)
}

func isPandigital(num string) bool {
	if len(num) != 9 {
		return false
	}
	for digit := 1; digit < 10; digit++ {
		if !strings.Contains(num, strconv.Itoa(digit)) {
			return false
		}
	}
	return true
}
