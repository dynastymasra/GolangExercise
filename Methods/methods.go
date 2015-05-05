package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (f MyFloat) AbsTwo() float64 {
	if f < 0 {
                return float64(-f)
        }
        return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println("methods")
	fmt.Println(v.Abs(), "\n")

	f := MyFloat(-math.Sqrt2)
	fmt.Println("methods continued")
	fmt.Println(f.AbsTwo(), "\n")

	v2 := &Vertex{3, 4}
	v2.Scale(5)
	fmt.Println("methods with pointer receiver")
	fmt.Println(v2, v2.Abs(), "\n")

	var a Abser
	b := MyFloat(-math.Sqrt2)
	n := &Vertex{3, 4}

	a = b
 	a = n

	fmt.Println("interfaces")
	fmt.Println(a.Abs(), "\n")
}
