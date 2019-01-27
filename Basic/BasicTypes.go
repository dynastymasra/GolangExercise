package main

import ("fmt"
	"math/cmplx"
	"math"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
var (
	Tobe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, Tobe, Tobe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var i int
	var l float64
	var b bool
	var s string = "test"
	fmt.Printf("Use percent v to show default value %v %v %v %v \n", i, l, b, s)
	fmt.Printf("Use percent q to show data type with value %q %q %q %q \n", i, l, b, s)
	fmt.Printf("Use percent T to show data type %T %T %T %T \n", i, l, b, s)

	var x, y int = 3, 4
	var a float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(a)
	fmt.Printf("Type conversion int %v %v %v, float64 %v \n", x, y, z, a)

	p, o, d, g := 42, 54.5, "test", true
	fmt.Printf("p is %T, o is %T, b is %T, g is %T \n", p, o, d, g)

	fmt.Println("This is test for constants")
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("Test", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
