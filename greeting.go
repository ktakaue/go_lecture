package main

import "fmt"
func main() {

// 小文字だとこのファイル内でしか使えない
	num := 1
// 大文字だと他のファイルからも使える
	NUM := 2


	fmt.Println(num, NUM)
}