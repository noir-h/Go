package statement

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func Defer() {
	// foo()
	// defer fmt.Println("world")

	// fmt.Println("hello")

	// defer is first in last out 
	/*
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
	*/

	file, _ := os.Open("./statement/for.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}