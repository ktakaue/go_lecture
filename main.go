package main

import (
	"fmt"
	"math/rand"
	"time"
)

func outer() {
	fmt.Println("time.Now()")
}

// HelloWorld
func main() {
	// console.log的な
	// timeをimportしてdate now()的な
	fmt.Println(time.Now())
	var hello string = "hello world Go";
	fmt.Println(hello)
	// 2ついれる
	var t,f bool = true, false
	fmt.Println(t)
	fmt.Println(f)
	var (
		a string = "A"
		b string = "B"
	)
	fmt.Println(a,b)
	var tetete  string = "kevin"
	fmt.Println(tetete)
	outer()
	fmt.Println(rand.Intn(10))
	pointer()
}


func anmoku()  {
	var i, j int = 1,2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i,j,k,c,python,java)

}

func kevin() {
	var sum  = 0
	for i :=0; i < 10; i++ {
		sum = i + 16
		fmt.Println(sum)
	}
}
func pointer(){
var point = "akasatana"
var resultPointer *string = &point
fmt.Println(resultPointer)
	}
