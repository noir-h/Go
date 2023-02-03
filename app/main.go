package main

import "fmt"

type Test struct {
	ID        int32
	firstName string
}

func main() {
	hoge := "hoge"
	if hoge == "hoge" {
		fmt.Println("hoge")
	} else if hoge == "hoge2" {
		fmt.Println("hoge2")
	} else if hoge == "hoge3" {
		fmt.Println("hoge3")
	}
	test(1, 2, 3, 4, 5, 6)
}

func test(a int, b int, c int, d int, e int, f int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
