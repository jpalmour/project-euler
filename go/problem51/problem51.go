package main

import (
	"fmt"
	"github.com/jpalmour/project-euler/go/primes"
	"strconv"
	"strings"
)

func main() {
	pc := primes.NewCalc(1000000)
	for _, prime := range pc.Primes {
		if has8PrimesViaReplacement(prime, pc) {
			break
		}
	}
}

func has8PrimesViaReplacement(num int, pc primes.Calc) bool {
	numStr := strconv.Itoa(num)
	for _, replacementStr := range getReplacementStrings(numStr) {
		primeCount := 0
		for repDigit := 1; repDigit < 10; repDigit++ {
			num := getRepNum(replacementStr, repDigit)
			if p, _ := pc.IsPrime(num); p {
				primeCount++
			}
		}
		if primeCount == 8 {
			// TODO: I origanally thought num was the solution, so I added this loop after the fact, making this a really weird solution
			for repDigit := 1; repDigit < 10; repDigit++ {
				num := getRepNum(replacementStr, repDigit)
				if p, _ := pc.IsPrime(num); p {
					fmt.Println(num)
					return true
				}
			}
		}
	}
	return false
}

func getReplacementStrings(num string) []string {
	repStrings := []string{}
	if len(num) == 0 {
		return repStrings
	}
	if len(num) == 1 {
		return []string{num, "*"}
	}
	for _, str := range getReplacementStrings(num[1:]) {
		a := string(num[0]) + str
		b := "*" + str
		repStrings = append(repStrings, a, b)
	}
	return repStrings
}

func getRepNum(repString string, repDigit int) int {
	num, _ := strconv.Atoi(strings.Replace(repString, "*", strconv.Itoa(repDigit), -1))
	return num
}
