package main

import "fmt"

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