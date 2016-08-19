package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cubesBySortedDigits := map[string][]int{}
	for num := 1; true; num++ {
		sortedDigits := sortDigits(num * num * num)
		cubesBySortedDigits[sortedDigits] = append(cubesBySortedDigits[sortedDigits], num*num*num)
		if len(cubesBySortedDigits[sortedDigits]) == 5 {
			fmt.Println(cubesBySortedDigits[sortedDigits][0])
			break
		}
	}
}

func sortDigits(num int) string {
	strSlice := strings.Split(strconv.Itoa(num), "")
	sort.Strings(strSlice)
	return strings.Join(strSlice, "")
}
