package main

import "fmt"

func main14() {
	x := make(map[string]int)
	x["angel"] = 10
	x["dhakal"] = 11
	for key, value := range x {
		fmt.Println(key, " => ", value)
	}
}
