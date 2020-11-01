package main

import "fmt"

func main() {
	a, b := calc(15, 10)
	fmt.Println(a, b)
	fmt.Println("Value Printed")
	x := []int{1, 2, 3, 4, 5}
	y := sliceAverage(x)
	fmt.Printf("average is %v\n", y)
}

func calc(x, y int) (addition int, subtraction int) {
	defer fmt.Println("x - y => ", x-y)
	defer fmt.Println("x + y => ", x+y)

	addition = x + y
	subtraction = x - y
	return
}

func sliceAverage(arr []int) (average int) {
	total := 0
	for _, value := range arr {
		total += value
	}
	average = total / len(arr)
	return
}
