package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int

	f1 := big.NewInt(0)
	f2 := big.NewInt(1)

	fmt.Print("Input f(n) value : ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i <= n; i++ {
		fmt.Printf("f(%d) = %d\n", i, f1)

		sum := big.NewInt(0)
		sum.Add(f1, f2)
		f1, f2 = f2, sum
	}
}
