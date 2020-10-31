package main

import (
	"fmt"
)

func main12() {
	x := []int{1, 2, 3, 4, 5}
	result := average(x)
	fmt.Println(result)
}

func average(arr []int) int {
	total := 0
	for _, i := range arr {
		total += i
	}
	return total
}
