package definition

import "fmt"

func foo(){
	// short valiables declaration
	// valid only in func
	xi := 1
	xf64 := 1.2
	xs := "tset"
	xt, xf := true, false 
	fmt.Println(xi, xf64, xs, xt, xf)
	// search of type
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
}
func Valiable(){
	// only declaration, the default initial value is set
	var(
		i int = 1
		f64 float64 = 1.2
		s string = "test" 
		t, f bool = true, false
	)
	fmt.Println(i, f64, s, t, f)

	foo()
}