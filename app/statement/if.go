package statement

import "fmt"

func by2(num int) string {
	if num % 2 == 0 {
		return "ok"
	}else {
		return "no"
	}
}

func If() {
	result := by2(4)
	if result == "ok" {
		fmt.Println("great")
	}
	// Can declaration
	fmt.Println(result)

	// 省略、ただifのスコープでしかresult2は使えない
	if result2 := by2(4); result2 == "ok" {
		fmt.Println("great")	
	}
	// Can not declaration
	// fmt.Println(result2)

	// num := 12
	// if num%2 == 0 {
	// 	fmt.Println("by 2")
	// } else if num%3 == 0 {
	// 	fmt.Println("by 3")
	// } else {
	// 	fmt.Println("else")
	// }

	// x, y := 10, 10
	// if x == 10 && y == 10 {
	// 	fmt.Println("&&")
	// }
	// if x == 10 || y == 10 {
	// 	fmt.Println("||")
	// }
}
