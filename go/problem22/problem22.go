package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Problem 22...")
	names := strings.Split(input, ",")
	sort.Strings(names)
	totalScore := 0
	for i, name := range names {
		totalScore += getScore(i+1, strings.Replace(name, "\"", "", -1))
	}
	fmt.Println(totalScore)
}

func getScore(position int, name string) int {
	return position * alphaValue(name)
}

func alphaValue(name string) int {
	total := 0
	for _, char := range name {
		total += int(char - 'A' + 1)
	}
	return total
}
