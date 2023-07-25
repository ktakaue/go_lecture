package main

import (
	"fmt"
	"reflect"
)
func main() {

	num := 123
	var num2 int = 456
	num3 := 1.23
	var num4 float64 = 4.5678911111
	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(num3))
	fmt.Println(reflect.TypeOf(num4))
}