package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// taking name as user input
	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	// taking age as user input
	fmt.Print("Enter your age: ")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Name is %v and age is %v\n", name, age)
}
