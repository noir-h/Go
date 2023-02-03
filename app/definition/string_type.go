package definition

import (
	"fmt"
	"strings"
)

func String_type() {
	fmt.Println("Hello World")            // Hello World
	fmt.Println("Hello " + "World")       // Hello World
	fmt.Println("Hello World"[0])         // 72 , アスキーコードになる
	fmt.Println(string("Hello World"[0])) // H

	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)                            // Xello World
	fmt.Println(strings.Contains(s, "World")) // true

	fmt.Println(`Test
	Test`)
	fmt.Println("\"")
	fmt.Println(`"`)
}
