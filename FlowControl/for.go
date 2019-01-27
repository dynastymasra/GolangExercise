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
	fmt.Println("for example")
	sum := 0
	for i :=0; i < 10; i++ {
		fmt.Print(" ", i)
		sum += i
	}
	fmt.Println("\n total =", sum)

	fmt.Println("\n for continued example")
	sum = 1
	for ; sum < 1000; {
		fmt.Print(" ", sum)
		sum += sum
	}
	fmt.Println("\n total =", sum)
}
