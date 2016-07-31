package main

import (
	"fmt"
)

var daysInMonthMap = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func main() {
	fmt.Println("Problem 19...")
	monthsStartingOnSunday := 0
	for year := 1901; year < 2001; year++ {
		for month := 1; month < 13; month++ {
			if monthStartsOnSunday(month, year) {
				monthsStartingOnSunday++
			}
		}
	}
	fmt.Println(monthsStartingOnSunday)
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func monthStartsOnSunday(month, year int) bool {
	return daysSince1Jan1900(month, year)%7 == 6
}

func daysSince1Jan1900(month, year int) int {
	days := 0
	for yearN := 1900; yearN < year; yearN++ {
		days += 365
		if isLeapYear(yearN) {
			days++
		}
	}
	for monthN := 1; monthN < month; monthN++ {
		days += daysInMonth(month, year)
	}
	return days
}

func daysInMonth(month, year int) int {
	if isLeapYear(year) && month == 2 {
		return 29
	}
	return daysInMonthMap[month]
}
