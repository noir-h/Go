package main

import "fmt"

type Test struct {
	ID        int32
	firstName string
}

func main() {

	mike := Test{ID: 1, firstName: "hoge"}
	if mike.firstName == "hoge" {
		fmt.Println("hoge")
	} else if mike.firstName == "hoge2" {
		fmt.Println("hoge2")
	} else if mike.firstName == "hoge3" {
		fmt.Println("hoge2")
	}
	test(1, 2, 3, 4, 5)
}

func test(a int, b int, c int, d int, e int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
