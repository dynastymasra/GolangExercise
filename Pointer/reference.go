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
