package main

import (
	"fmt"
	"math"
)

func main() {
	triangleMap := getMap(func(x int) float64 {
		return (math.Sqrt(8*float64(x)+1) - 1) / 2
	})
	squareMap := getMap(func(x int) float64 {
		return math.Sqrt(float64(x))
	})
	pentagonMap := getMap(func(x int) float64 {
		return (math.Sqrt(24*float64(x)+1) + 1) / 6
	})
	hexagonMap := getMap(func(x int) float64 {
		return (math.Sqrt(8*float64(x)+1) + 1) / 4
	})
	heptagonMap := getMap(func(x int) float64 {
		return (math.Sqrt(40*float64(x)+9) + 3) / 10
	})
	octagonMap := getMap(func(x int) float64 {
		return (math.Sqrt(3*float64(x)+1) + 1) / 3
	})
	set := getSets([]map[int][]int{triangleMap, squareMap, pentagonMap, hexagonMap, heptagonMap, octagonMap})
	sum := 0
	for _, num := range set[0] {
		sum += num
	}
	fmt.Println(sum)
}

func getSets(unvisited []map[int][]int) [][]int {
	result := [][]int{}
	for first2, last2s := range unvisited[0] {
		for _, last2 := range last2s {
			num := first2*100 + last2
			result = append(result, getSetsRecur(unvisited[0], last2, []int{num}, unvisited[1:])...)
		}
	}
	return result
}

func getSetsRecur(start map[int][]int, next int, set []int, unvisited []map[int][]int) [][]int {
	result := [][]int{}
	if len(unvisited) == 0 {
		if next == set[0]/100 {
			return [][]int{set}
		}
		return [][]int{}
	}
	for i := 0; i < len(unvisited); i++ {
		for _, last2 := range unvisited[i][next] {
			newSet := append(set, next*100+last2)
			dup := make([]map[int][]int, len(unvisited))
			copy(dup, unvisited)
			newUnvisited := append(dup[0:i], dup[i+1:]...)
			result = append(result, getSetsRecur(start, last2, newSet, newUnvisited)...)
		}
	}
	return result
}

func getMap(f func(int) float64) map[int][]int {
	result := make(map[int][]int)
	for x := 1000; x < 10000; x++ {
		n := f(x)
		if float64(int(n)) == n {
			result[x/100] = append(result[x/100], x%100)
		}
	}
	return result
}
