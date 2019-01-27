package main

import (
  "fmt"
  "math"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
func findDigit(d []int) (int, int, int)  {
  positive := 0
  negative := 0
  zero := 0
  for i := 0; i < len(d); i++ {
    if d[i] > 0 {
      positive++
    } else if d[i] < 0 {
      negative++
    } else {
      zero++
    }
  }

  return positive, negative, zero
}

func divide(positive int, negative int, zero int, n int) (float64, float64, float64)  {
  var positiveDiv, negativeDiv, zeroDiv float64
  positiveDiv = float64(positive)/float64(n)
  negativeDiv = float64(negative)/float64(n)
  zeroDiv = float64(zero)/float64(n)

  return positiveDiv, negativeDiv, zeroDiv
}

func main() {
  var d []int
  var n int
  var v int

  fmt.Print("Input Max Array n > ")
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    fmt.Printf("Input Array Index |%v| > ", i)
    fmt.Scan(&v)
    d = append(d, v)
  }
  fmt.Println("Array > ", d)
  positive, negative, zero := findDigit(d)
  fmt.Printf("Positive[%v] Negative[%v] Zero[%v]\n", positive, negative, zero)
  q, w, e := divide(positive, negative, zero, n)
  fmt.Printf("Value +[%v] -[%v] 0[%v]\n", round(q, .5, 3), round(w, .5, 3), round(e, .5, 3))
}

func round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
