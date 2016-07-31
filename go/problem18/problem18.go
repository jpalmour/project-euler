package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Problem 18...")
	lines := strings.Split(input, "\n")
	rows := [][]int{}
	for i, line := range lines {
		rows = append(rows, []int{})
		nums := strings.Split(line, " ")
		for j, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			maxPath := num + maxParent(i, j, rows)
			rows[i] = append(rows[i], maxPath)
		}
	}
	fmt.Println(max(rows[len(rows)-1]))
}

func maxParent(i, j int, tree [][]int) int {
	left := 0
	right := 0
	if i-1 >= 0 {
		if j-1 >= 0 {
			left = tree[i-1][j-1]
		}
		if j < len(tree[i-1]) {
			right = tree[i-1][j]
		}
	}
	if left > right {
		return left
	}
	return right
}

func max(rows []int) int {
	max := 0
	for _, j := range rows {
		if j > max {
			max = j
		}
	}
	return max
}
