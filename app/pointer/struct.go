package pointer

import "fmt"

type Vertex struct{
	/*
	X int
	Y int
	*/
	X, Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	// goは自動でアドレスの実態を見に行く
	v.X = 1000
}


func Struct() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)

	// v := Vertex{X: 1, Y:2}
	// fmt.Println(v)
	// fmt.Println(v.X, v.Y)
	// v.X = 100
	// fmt.Println(v.X, v.Y)

	// v2 := Vertex{X: 1}
	// fmt.Println(v2)

	// v3 := Vertex{1, 2, "test"}
	// fmt.Println(v3)

	// v4 := Vertex{}
	// fmt.Printf("%T %v\n", v4, v4)

	// map, sliceと違いnilにならない
	var v5  Vertex
	fmt.Printf("%T %v\n", v5, v5)
	/* 
	** example for slice and map
	var s []int
  var m map[string]int
	if s == nil {
		fmt.Println("nil")
	}
	if m == nil {
		fmt.Println("nil")
	}
	*/

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

	s := make([]int, 0) // 推奨
	// s := []int{}
	fmt.Println(s)
}