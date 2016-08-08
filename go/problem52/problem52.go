package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for num := 1; true; num++ {
		digitString := getDigitsString(num)
		match := true
		for n := 2; n < 7; n++ {
			if digitString != getDigitsString(num*n) {
				match = false
			}
		}
		if match {
			fmt.Println(num)
			break
		}
	}
}

func getDigitsString(num int) string {
	str := strconv.Itoa(num)
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
