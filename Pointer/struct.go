package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	pOne = &Vertex{1, 2}
)

func main() {
	fmt.Println("struct")
	fmt.Println(Vertex{1, 2})

	fmt.Println("struct fields")
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	fmt.Println("struct pointer")
	v = Vertex{1, 2}
	p := &v
	p.X = 1000
	fmt.Println(v)

	fmt.Println("struct literal")
	fmt.Println(v1, pOne, v2, v3)
}
	
