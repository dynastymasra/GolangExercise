package main

import (
  "fmt"
  "math/big"
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
  var n big.Int

  fmt.Print("Input Value Big n : ")
  fmt.Scan(&n)
  fmt.Println("Factorial Big >", factorial(&n))
  fmt.Println("Value Factorial Dynamic 10")
  r := big.NewInt(10)
  fmt.Println("Factorial >", factorial(r))
}

func factorial(n *big.Int) (result *big.Int) {
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		fmt.Println("n = ", n)
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}
