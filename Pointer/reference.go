package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Value is", x)
	fmt.Println("Address value is", &x)
	changeValueNow(&x)
}

func changeValueNow(x *int) {
	*x = 2
	fmt.Println("Value func is", *x) 
	fmt.Println("Address func is", x)
}