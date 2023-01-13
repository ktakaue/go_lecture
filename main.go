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
	var kevin string = "kevin";
	fmt.Println(kevin)
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
}

func kevin(x, y int) int {
	return x + y
}

func anmoku()  {
	var i, j int = 1,2
	k := 3
	var c, python, java = true, false, "no!"
	fmt.Println(i,j,k,c,python,java)

}