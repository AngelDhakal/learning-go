package main

import (
	"fmt"
)

func main10() {
	var x = [5]int{10, 10, 10, 10, 10}
	total := 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total)
}
