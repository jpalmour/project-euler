package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	value int
	suit  byte
}

type hand []card

func (h hand) Len() int {
	return len(h)
}

func (h hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hand) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func main() {
	handsWon := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hand1 := hand([]card{})
		hand2 := hand([]card{})
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
	val, _ := strconv.Atoi(str[0:1])
	if val == 0 {
		toVal := map[string]int{"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
		val = toVal[str[0:1]]
	}
	x := card{value: val, suit: str[1]}
	return x
}

func hand1Wins(hand1, hand2 hand) bool {
	sort.Sort(hand1)
	sort.Sort(hand2)
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

func getValue(h hand) (rank, rankValue, highestCard int) {
	if r, rV, hC, isStraightFlush := straightFlushCheck(h); isStraightFlush {
		return r, rV, hC
	}
	if r, rV, hC, isFourOfAKind := fourOfAKindCheck(h); isFourOfAKind {
		return r, rV, hC
	}
	if r, rV, hC, isFullHouse := fullHouseCheck(h); isFullHouse {
		return r, rV, hC
	}
	if r, rV, hC, isFlush := flushCheck(h); isFlush {
		return r, rV, hC
	}
	if r, rV, hC, isStraight := straightCheck(h); isStraight {
		return r, rV, hC
	}
	if r, rV, hC, isThreeOfAKind := threeOfAKindCheck(h); isThreeOfAKind {
		return r, rV, hC
	}
	if r, rV, hC, isTwoPairs := twoPairsCheck(h); isTwoPairs {
		return r, rV, hC
	}
	if r, rV, hC, isOnePair := onePairCheck(h); isOnePair {
		return r, rV, hC
	}
	return highCardValue(h)
}

const (
	HIGH_CARD       = iota
	PAIR            = iota
	TWO_PAIRS       = iota
	THREE_OF_A_KIND = iota
	STRAIGHT        = iota
	FLUSH           = iota
	FULL_HOUSE      = iota
	FOUR_OF_A_KIND  = iota
	STRAIGHT_FLUSH  = iota
)

func valueMap(h hand) map[int]int {
	vm := map[int]int{}
	for _, c := range h {
		vm[c.value]++
	}
	return vm
}

func straightFlushCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	for i := 1; i < len(h); i++ {
		if h[i].value != h[i-1].value+1 || h[i].suit != h[i-1].suit {
			return 0, 0, 0, false
		}
	}
	return STRAIGHT_FLUSH, h[4].value, h[4].value, true
}

func fourOfAKindCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	for val, count := range valueMap(h) {
		if count == 4 {
			return FOUR_OF_A_KIND, val, h[4].value, true
		}
	}
	return 0, 0, 0, false
}

func fullHouseCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	threeOfAKindValue := 0
	foundPair := false
	for val, count := range valueMap(h) {
		if count == 3 {
			threeOfAKindValue = val
		}
		if count == 2 {
			foundPair = true
		}
	}
	if threeOfAKindValue != 0 && foundPair {
		return FULL_HOUSE, threeOfAKindValue, h[4].value, true
	}
	return 0, 0, 0, false
}

func flushCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	for i := 1; i < len(h); i++ {
		if h[i].suit != h[i-1].suit {
			return 0, 0, 0, false
		}
	}
	return FLUSH, h[4].value, h[4].value, true
}

func straightCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	for i := 1; i < len(h); i++ {
		if h[i].value != h[i-1].value+1 {
			return 0, 0, 0, false
		}
	}
	return STRAIGHT, h[4].value, h[4].value, true
}

func threeOfAKindCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	for val, count := range valueMap(h) {
		if count == 3 {
			return THREE_OF_A_KIND, val, h[4].value, true
		}
	}
	return 0, 0, 0, false
}

func twoPairsCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	pairValues := []int{}
	for val, count := range valueMap(h) {
		if count == 2 {
			pairValues = append(pairValues, val)
		}
	}
	if len(pairValues) == 2 {
		bestPair := pairValues[0]
		if pairValues[1] > pairValues[0] {
			bestPair = pairValues[1]
		}
		return TWO_PAIRS, bestPair, h[4].value, true
	}
	return 0, 0, 0, false
}

func onePairCheck(h hand) (rank, rankValue, highestCard int, match bool) {
	for val, count := range valueMap(h) {
		if count == 2 {
			return PAIR, val, h[4].value, true
		}
	}
	return 0, 0, 0, false
}

func highCardValue(h hand) (rank, rankValue, highestCard int) {
	return HIGH_CARD, h[4].value, h[4].value
}
