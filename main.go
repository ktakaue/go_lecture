package main

func main(){
	// println("Hello, World!")
	// print(time.Now().String())

	// 適当に変数を定義する
	var a int = 1
	b  := 2
	var c, d int = 3, 4
	println(a, b, c, d)

	e,f,g,h := 5,6,7,8
	// 改行しながら出力
	println(e)
	println(f)
	println(g)
	println(h)

	// 配列
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	println(arr[0], arr[1], arr[2])

	// 配列の初期化
	arr2 := [3]int{1, 2, 3}
	println(arr2[0], arr2[1], arr2[2])

	// 配列の長さを取得
	arr3 := [...]int{444, 555, 666, 777, 888, 999}
	println((len(arr3)))


}