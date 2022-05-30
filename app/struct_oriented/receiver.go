package structoriented

import "fmt"

type Vertex struct{
	x, y int
}

// value reciever
func (v Vertex) Area() int {
	return v.x * v.y
}

// pointa reciever
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func Area(v Vertex) int {
	return v.x * v.y
}

// Embedded
type Vertex3D struct{
	// Vertexを埋め込んでる
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// constoructor
// func New(x, y int) * Vertex {
// 	return &Vertex{x, y}
// }
func New(x, y, z int) * Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func Receiver() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(Area(v))

// constoructor
	c := New(3, 4, 5)
	c.Scale3D(10)
	// fmt.Println(c.Area())
	fmt.Println(c.Area3D())
}