package pointer

import "fmt"

func New() {
	/*
	var p *int = new(int)
	fmt.Println(*p)	
	*p ++
	fmt.Println(*p)	
	*/

	/*
	var p2 *int 
	fmt.Println(p2)	
	*p2 ++
	fmt.Println(*p2)	
	*/

	// makeとnewの違い、newはポインタを返す(chanel, struct etc)
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p *int = new(int)
	fmt.Printf("%T\n", p)

}