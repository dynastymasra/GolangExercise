package main

import "fmt"

func main() {
	c := make(chan int, 4)
	c <- 5
	c <- 4
	c <- 6
	c <- 2
	fmt.Println(<-c, <-c, <-c, <-c)
}