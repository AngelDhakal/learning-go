package main

import "fmt"

func main6() {
	fmt.Print("Enter first number: ")
	var input1, input2 int
	fmt.Scanf("%d", &input1)
	fmt.Print("Enter second number: ")
	fmt.Scanf("%d", &input2)
	output := input1 + input2
	println(output)
}
