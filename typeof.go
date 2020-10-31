package main

import (
	"fmt"
	"reflect"
)

func main16() {
	x := []int{1, 2, 3}
	y := [3]int{1, 2, 3}
	z := "angel"
	aa := 10
	ab := 10.10
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(reflect.TypeOf(aa))
	fmt.Println(reflect.TypeOf(ab))
}
