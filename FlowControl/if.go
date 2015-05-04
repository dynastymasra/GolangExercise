package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func powAgain(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
	

func main() {
	fmt.Println("if example")
	fmt.Println(sqrt(4), sqrt(-4))
	
	fmt.Println("\nif short statement example")
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println("\nif else example")
	fmt.Println(powAgain(3, 2, 10), powAgain(3, 3, 20))
}
