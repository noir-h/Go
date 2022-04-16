package main

import (
	"app/how_to_import"
	"fmt"
	"os/user"
)

/*
func init(){
	fmt.Println("init")
}
*/

func fizz(){
	fmt.Println("fizz")
}
func main() {
	// fizz()
	how_to_import.Test()
	fmt.Println(user.Current())
	// fmt.Println("Hello golang from docker!")
}

// how to read doc
// $ go doc package name  ex) go doc package fmt Println