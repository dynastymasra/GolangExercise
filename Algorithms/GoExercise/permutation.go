package main

import "fmt"

func main() {
	var n int

	fmt.Print("Input f(n) value : ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []string{"A", "B", "C"}

	printString(data, "", n, n)
}

func printString(data []string, res string, n, v int) {
	if v == 0 {
		fmt.Println(res)
		return
	}

	for i := 0; i < n; i++ {
		newRes := res + data[i]
		printString(data, newRes, n, v-1)
	}
}
