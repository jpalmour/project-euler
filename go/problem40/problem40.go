package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Problem 40...")
	d := 1
	searchingFor := 1
	product := 1
	for n := 1; d < 1000000; n++ {
		current := strconv.Itoa(n)
		l := len(current)
		d += l
		if searchingFor >= d-l && searchingFor < d {
			index := searchingFor - d + l
			val, _ := strconv.Atoi(current[index : index+1])
			product *= val
			searchingFor *= 10
		}
	}
	fmt.Println(product)
}
