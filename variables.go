package main

import (
	"fmt"
	"strconv"
)

var x int = 10
var y int = 10

func main5() {
	fmt.Println(x == y)
	f()
}

func f() {
	fmt.Println("x is " + strconv.Itoa(x) + " y is " + strconv.Itoa(y))
}
