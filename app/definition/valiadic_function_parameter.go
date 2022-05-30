package definition

import "fmt"

func foo_y(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func Valiadic_function() {
	foo_y()
	foo_y(10, 20)
	foo_y(10, 20, 30)

	s := []int{10, 20, 30}
	fmt.Println(s)

	foo_y(s...)
}
