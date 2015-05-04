package main

import "fmt"

var c, python, java = true, false, "past"

func main() {
	i, j := 1, 2
	var x int = i + j
	fmt.Println(i, j, x, c, python, java)
}

