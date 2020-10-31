package main

import "fmt"

func main15() {
	x := [5]int{1, 2, 3, 4, 5}
	y := x[1:3]
	z := y[:cap(y)]
	fmt.Println(z)
}
