package main

import "fmt"

func main9() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " Even")
		} else if i%2 == 1 {
			fmt.Println(i, " Odd")
		}
	}
}
