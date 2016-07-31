package main

import (
	"fmt"
)

func main() {
	fmt.Println("Problem 21...")
	perms := perms("0123456789")
	fmt.Println(perms[999999])
}

func perms(str string) []string {
	if str == "" {
		return []string{str}
	}
	result := []string{}
	for i := 0; i < len(str); i++ {
		recur := str[:i] + str[i+1:]
		result = append(result, appendToFront(str[i:i+1], perms(recur))...)
	}
	return result
}

func appendToFront(head string, strs []string) []string {
	for i, str := range strs {
		strs[i] = head + str
	}
	return strs
}
