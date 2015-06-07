package main

import "fmt"

type incFuncType func()
type getFuncType func() int
	
func myCounter(initCount int) (incFuncType, getFuncType) {
	count := initCount
	incriment := func() {
		count++
		fmt.Println("Count is", count)
	}
	get := func() int {
		fmt.Println("Final count is", count)
		return count
	}

	return incriment, get
}

func main() {
	incriment, get := myCounter(10)
	for i:=0; i<5; i++ {
		incriment()
	}
	fmt.Println("Final value is", get())
}