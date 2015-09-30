package main

import (
	"fmt"
	"time"
	"math/rand"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(10)
	if random < 5 {
		fmt.Println("Number is < 5")
	}
	fmt.Println("Time Now :", time.Now())
	fmt.Println("Time Now UTC :", time.Now().UTC())
	fmt.Println("Time Now UTC UnixNano :", time.Now().UTC().UnixNano())
	fmt.Println("Number :", random)
}
