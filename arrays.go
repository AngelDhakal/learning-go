package main

import (
	"fmt"
)

func main10() {
	var x = [5]int{10, 10, 10, 10, 10}
	total := 0
	for index, value := range x {
		total += value
		Use(index)
	}
	fmt.Println(total)
}

func Use(vals ...interface{}) {
	for _, value := range vals {
		_ = value
	}
}
