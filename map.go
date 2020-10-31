package main

import "fmt"

func main17() {

	// first way of making map
	fruits := map[string]int{
		"apple":  10,
		"orange": 5,
		"mango":  10,
	}

	// second way of creating map
	meats := make(map[string]int)

	// adding values to a map
	meats["chicken"] = 5
	meats["mutton"] = 2
	meats["steak"] = 2
	meats["beef"] = 4

	// accessing a value in a map
	fmt.Println(fruits["apple"])
	fmt.Println(meats["chicken"])

	// changing a value in a map
	fruits["apple"] = 15
	meats["chicken"] = 4

	// deleting a value in a map
	delete(fruits, "mango")
	delete(meats, "steak")
	delete(meats, "beef")

	fmt.Println(fruits)
	fmt.Println(meats)

	// length of maps
	fmt.Println(len(fruits))
	fmt.Println(len(meats))

	// checks if specific value exists in a map
	fmt.Println(isKeyPresent(fruits, "mango"))
}

// function to check whether a key exists in a map or not

func isKeyPresent(a map[string]int, b string) bool {
	_, ok := a[b]
	if ok == true {
		return true
	}
	return false
}
