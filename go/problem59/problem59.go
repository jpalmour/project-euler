package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	msgInts := []int{}
	for _, tokens := range strings.Split(input, ",") {
		intVal, _ := strconv.Atoi(tokens)
		msgInts = append(msgInts, intVal)
	}
Loop:
	for x := 97; x <= 122; x++ {
		for y := 97; y <= 122; y++ {
			for z := 97; z <= 122; z++ {
				decodedMsg, sum := decode(msgInts, []int{x, y, z})
				if probablyEnglish(decodedMsg) {
					fmt.Println(sum)
					break Loop
				}
			}
		}
	}
}

func decode(encodedMsg, key []int) (string, int) {
	msg := ""
	sum := 0
	for i, asciiCode := range encodedMsg {
		xorVal := key[i%len(key)]
		decodedVal := asciiCode ^ xorVal
		msg += string(decodedVal)
		sum += decodedVal
	}
	return msg, sum
}

func probablyEnglish(str string) bool {
	return strings.Contains(str, "the") &&
		strings.Contains(str, "be") &&
		strings.Contains(str, "to") &&
		strings.Contains(str, "of") &&
		strings.Contains(str, "and")
}
