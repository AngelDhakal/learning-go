package main

import (
	"fmt"
)


/*
Functions
 - they are their own data types
 - they can be assigned to variables
 - they can be passed as function parameters
 - they can be returned by another function
*/

func main() {
	a, b := calc(15, 10)
	fmt.Println(a, b)
	fmt.Println("Value Printed")
	x := []int{1, 2, 3, 4, 5}
	y := sliceAverage(x)
	fmt.Printf("average is %v\n", y)

	// function to add two numbers
	// the scope of this function is within main function only
	add := func(a, b int) int {
		return a + b
	}(20,10)
	fmt.Printf("Addition of two numbers are %v \n", add)
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
