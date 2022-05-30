package structoriented

import (
	"fmt"
)

func do(i interface{}) {
	// interface -> int 等の型変換はタイプアサーション
	// ii := i.(int)

	// switchとtypeはセット
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func TypeAssertion() {
	do(10)
	do("Mike")
	do(true)

	// キャストはこっち
	var i int = 10
	ii := float64(i)
	fmt.Println(i,ii)
}