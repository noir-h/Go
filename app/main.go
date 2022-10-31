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
}
