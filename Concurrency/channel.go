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
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total //send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x + y)
}
