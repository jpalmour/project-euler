package main

import (
	"fmt"
	"strconv"
	"strings"
)

type card struct {
	value int
	suit  byte
}

func main() {
	handsWon := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hand1 := []card{}
		hand2 := []card{}
		for i, card := range strings.Split(line, " ") {
			if i < 5 {
				hand1 = append(hand1, getCard(card))
			} else {
				hand2 = append(hand2, getCard(card))
			}
		}
		if hand1Wins(hand1, hand2) {
			handsWon++
		}
	}
	fmt.Println(handsWon)

}

func getCard(str string) card {
	val, _ := strconv.Atoi(string(str[0:1]))
	x := card{value: val, suit: str[1]}
	return x
}

func hand1Wins(hand1, hand2 []card) bool {
	rank1, rankValue1, highestCard1 := getValue(hand1)
	rank2, rankValue2, highestCard2 := getValue(hand2)
	if rank1 == rank2 {
		if rankValue1 == rankValue2 {
			return highestCard1 > highestCard2
		}
		return rankValue1 > rankValue2
	}
	return rank1 > rank2
}

func getValue(hand []card) (rank, rankValue, highestCard int) {

}
