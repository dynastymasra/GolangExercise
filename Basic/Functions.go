package main

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Simple example function in Go")
	fmt.Println("Simple function in go add =", add(40, 30))

	a, b := swap("hello", "world")
	fmt.Print("Simple fucntion swap = ")
	fmt.Println(a, b)

	fmt.Print("Function with naked return = ")
	fmt.Println(split(17))
}
