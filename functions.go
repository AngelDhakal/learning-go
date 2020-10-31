package main

import "fmt"

func main() {
	a, b := calc(15, 10)
	fmt.Println(a, b)
	fmt.Println("Value Printed")

}

func calc(x, y int) (addition int, subtraction int) {
	defer fmt.Println("x - y => ", x-y)
	defer fmt.Println("x + y => ", x+y)

	addition = x + y
	subtraction = x - y
	return
}
