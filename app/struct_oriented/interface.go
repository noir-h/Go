package structoriented

import "fmt"

type Human interface {
	say() string
}

type Person struct {
	Name string
}

// implの強制
func (p *Person) say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

func DriveCar(human Human) {
	if human.say() == "Mr.sora" {
		fmt.Println("run")
	} else {
		fmt.Println("Get out")
	}
}

func Interface() {
	var sora Human = &Person{"sora"}
	DriveCar(sora)
}