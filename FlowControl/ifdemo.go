package main

import (
	"fmt"
	"time"
	"math/rand"
)

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