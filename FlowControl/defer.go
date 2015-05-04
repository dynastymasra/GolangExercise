package main

import "fmt"

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
