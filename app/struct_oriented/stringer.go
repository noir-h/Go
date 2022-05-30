package structoriented

import "fmt"

type person struct {
	Name string
	age int
}

// printlnした時に出力内容を上書きしたい時に使う
func (p person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func Stringer() {
	sora := person{"sora", 24}
	fmt.Println(sora)
}