package main

import "fmt"

func main() {

	// 配列の要素数を指定せずに宣言する
	a :=[...]string{"hello","world","!"}
	fmt.Println(a[0],a[1],a[2])


	// スライスで宣言する
	b :=[]string{"hello","world","!"}
	b = append(b, "appendkevikevi")
	fmt.Println(b[0],b[1],b[2],b[3])
}