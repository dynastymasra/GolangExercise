package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int

	fmt.Print("Input f(n) value : ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	f := make([]*big.Int, n+1)
	f[0] = big.NewInt(1)

	fmt.Printf("f(%d) = %d\n", 0, f[0])

	for i := 1; i <= n; i++ {
		f[i] = big.NewInt(0)
		for n := 1; n <= 6; n++ {
			if (i - n) >= 0 {
				f[i].Add(f[i], f[i-n])
			}
		}
		fmt.Printf("f(%d) = %d\n", i, f[i])
	}
}
