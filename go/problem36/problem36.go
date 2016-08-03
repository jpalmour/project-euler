package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Problem 36...")
	sum := 0
	for num := 1; num < 1000000; num++ {
		bin := strconv.FormatInt(int64(num), 2)
		dec := fmt.Sprintf("%v", num)
		if isPalindrome(bin) && isPalindrome(dec) {
			sum += num
		}
	}
	fmt.Println(sum)
}

func isPalindrome(str string) bool {
	return reverse(str) == str
}

func reverse(str string) string {
	result := ""
	for _, char := range str {
		result = string(char) + result
	}
	return result
}
