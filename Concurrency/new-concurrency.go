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
func loop(start, end int) int {
	var i int
	for k := start; k < end; k++ {
		i = k
	}
	return i
}

func main() {
	p := make(chan int)

	go func() {
		t := loop(100, 222)
		p <- t
	}()

	go func() {
		t := loop(234, 455)
		p <- t
	}()

	z := <-p
	v := <-p
	fmt.Println(v, z)
}
