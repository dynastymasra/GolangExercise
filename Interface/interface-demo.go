package main

import "fmt"

type animal interface {
	canEat(food string) bool
}

type duck struct {	
	canSwim bool
}

func (a duck) canEat(food string) bool {
	if food == "fish" {
		return true
	}
	return false
}

func main() {
	var myDuck animal
	myDuck = duck{true}
	fmt.Println("Duck eat fish?", myDuck.canEat("fish"))
	fmt.Println("Duck can swim?", myDuck)
}