package definition

import "fmt"

func Add(x int, y int) {
	fmt.Println("add function")
	fmt.Println(x + y)
}

func Addx(x int, y int) int {
	return x + y
}

func Addy(x, y int) (int, int) {
	return x + y, x * y
}

func Cal(price, item int) (result int) {
	result = price * item
	// return result
	// 返り値を宣言している場合はreturnだけでok
	return
}
