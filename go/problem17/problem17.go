package main

import (
	"fmt"
)

var (
	units = map[int]int{
		0: 0,
		1: 3,
		2: 3,
		3: 5,
		4: 4,
		5: 4,
		6: 3,
		7: 5,
		8: 5,
		9: 4,
	}
	ten = map[int]int{
		0: 0,
		1: 0,
		2: 6,
		3: 6,
		4: 5,
		5: 5,
		6: 5,
		7: 7,
		8: 6,
		9: 6,
	}
	teens = map[int]int{
		0: 3,
		1: 6,
		2: 6,
		3: 8,
		4: 8,
		5: 7,
		6: 7,
		7: 9,
		8: 8,
		9: 8,
	}
)

func main() {
	fmt.Println("Problem 17...")
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += numToWordsLength(i)
	}
	fmt.Println(sum)
}

func numToWordsLength(num int) int {
	if num == 1000 {
		return 11
	}
	return getHundredsLength(num) + getTensLength(num) + getUnitsLength(num)
}

func getHundredsLength(num int) int {
	hundredsPlace := num / 100
	if hundredsPlace == 0 {
		return 0
	}
	hundredsValue := units[hundredsPlace]
	if num%100 != 0 {
		hundredsValue += 3
	}
	return hundredsValue + 7
}

func getTensLength(num int) int {
	tensPlace := (num % 100) / 10
	return ten[tensPlace]
}

func getUnitsLength(num int) int {
	tensPlace := (num % 100) / 10
	unitsPlace := num % 10
	if tensPlace == 1 {
		return teens[unitsPlace]
	}
	return units[unitsPlace]
}
