package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Problem 42...")
	names := strings.Split(input, ",")
	count := 0
	for _, name := range names {
		if isTriangleName(strings.Replace(name, "\"", "", -1)) {
			count++
		}
	}
	fmt.Println(count)
}

func isTriangleName(name string) bool {
	return isTriangleNumber(alphaValue(name))
}

func isTriangleNumber(num int) bool {
	for n := 1; num >= (n*(n+1))/2; n++ {
		if num == (n*(n+1))/2 {
			return true
		}
	}
	return false
}

func alphaValue(name string) int {
	total := 0
	for _, char := range name {
		total += int(char - 'A' + 1)
	}
	return total
}
