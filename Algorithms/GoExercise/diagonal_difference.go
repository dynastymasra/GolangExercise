package main

import (
  "fmt"
  "math"
)

func printMatrixs(d []int, n int)  {
  min := 0
  max := n
  for i := 0; i < n; i++ {
    fmt.Println(d[min:max])
    min = max
    max += n
  }
}

func sumLeft(d []int, n int) int {
  min := 0
  max := n
  total := 0
  for i := 0; i < n; i++ {
    v := d[min:max]
    total += v[i]
    min = max
    max += n
  }

  return total
}

func sumRight(d []int, n int) int {
  min := 0
  max := n
  total := 0
  for i := n-1; i >= 0; i-- {
    v := d[min:max]
    total += v[i]
    min = max
    max += n
  }

  return total
}

func different(a int, b int) float64 {
  return math.Abs(float64(a-b))
}

func main() {
  var n int
  var v int
  var m []int

  fmt.Print("Input Max n > ")
  fmt.Scan(&n)
  for y := 0; y < n; y++ {
    for x := 0; x < n; x++ {
      fmt.Printf("Input Variable |%v %v| > ", x, y)
      fmt.Scan(&v)
      m = append(m, v)
    }
  }
  printMatrixs(m, n)
  fmt.Println("Total Sum Left > ", sumLeft(m, n))
  fmt.Println("Total Sum Right > ", sumRight(m, n))
  fmt.Println("Total Sum Diagonal > ", different(sumLeft(m, n), sumRight(m, n)))
}
