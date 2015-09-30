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
func main() {
	defer fmt.Println("\nworld")
	fmt.Println("hello")

	fmt.Println()
	fmt.Println("stacking defer")
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Print(" ", i)
	}
	fmt.Println("done")
}
