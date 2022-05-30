package main

import (
	structoriented "app/struct_oriented"
	"fmt"
)

/*
func init(){
	fmt.Println("init")
}
*/

func fizz() {
	fmt.Println("fizz")
}
func main() {
	// cast_of_type.Test()

	// definition.How_to_import()
	// definition.Const()
	// definition.Valiable()
	// definition.Int_type()
	// definition.String_type()
	// definition.Boolean_type()
	// definition.Cast_of_type()
	// definition.Array()
	// definition.Slice()
	// definition.Make_cap()
	// definition.Map()
	// definition.Byte()

	/*
	definition.Add(3, 5)
	r := definition.Addx(3, 5)
	fmt.Println(r)
	r1, r2 := definition.Addy(3, 5)
	fmt.Println(r1, r2)
	r3 := definition.Cal(5, 6)
	fmt.Println(r3)
	// 無名
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
	// 即時
	func(x int) {
		fmt.Println("inner func", x)
	}(2)
	*/
	
	// definition.Closure()
	// definition.Valiadic_function()
	// statement.If()
	// statement.For()
	// statement.Range()
	// statement.Switch()
	// statement.Defer()
	// statement.Log()
	// statement.Error()
	// statement.Panic()
	// pointer.Pointa()
	// pointer.New()
	// pointer.Struct()
	// structoriented.Receiver()
	// structoriented.Interface()
	// structoriented.TypeAssertion()
	// structoriented.Stringer()
	structoriented.CustomError()
}

// how to read doc
// $ go doc package name  ex) go doc package fmt Println
