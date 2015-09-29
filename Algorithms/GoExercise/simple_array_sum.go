package main

import "fmt"

func main() {
  var n int
  var v []int
  var a int
  var total int

  fmt.Print("Input Total Array n > ")
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    fmt.Print("Input Index Array ", i, " > ")
    fmt.Scan(&a)
    v = append(v, a)
  }

  for _, i := range v {
    total += i
  }

  fmt.Println("Total > ", total)
}
