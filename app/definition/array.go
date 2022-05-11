package definition

import "fmt"

func Array() {
	var a [2]int
	a[0] = 100
	a[1] = 100
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	// Cannot resize array
	// b = append(b, 300)
	fmt.Println(b)
}
