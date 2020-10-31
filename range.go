package main

import "fmt"

func main14() {
	x := []int{1, 2, 3, 4, 5, 2, 7, 8, 9}

	for i, value := range x {
		for j := i + 1; j < len(x); j++ {
			value2 := x[j]
			if value == value2 {
				fmt.Println(value)
			}
		}
	}

}
