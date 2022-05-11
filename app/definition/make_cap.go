package definition

import "fmt"

func Make_cap() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%d\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%d\n", len(n), cap(n), n)
	// sliceにおいてはcapを超えるサイズを入れることが可能
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%d\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%d\n", len(a), cap(a), a)

	// メモリ確保
	b := make([]int, 0)
	// メモリ確保なし
	var c []int
	fmt.Printf("len=%d cap=%d value=%d\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%d\n", len(c), cap(c), c)

	// valueを確保
	// c = make([]int, 5)
	// valueは確保していない
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
