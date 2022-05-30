package structoriented

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T %v" , i, i)
	fmt.Printf("%T %v" , 1, 1)
	// iはMyInt型だが返り値の型がint型なのでキャスト
	return int(i * 2)
}

func NonStruct() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}