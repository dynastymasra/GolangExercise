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
	c := make(chan int, 4)
	c <- 5
	c <- 4
	c <- 6
	c <- 2
	fmt.Println(<-c, <-c, <-c, <-c)
}
