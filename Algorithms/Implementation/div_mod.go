package main

import "fmt"

func main() {
  var n int
  var v int

  fmt.Println("Input Value n :")
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    v = 0
    fmt.Println("Input Value :")
    fmt.Scan(&v)

    m := v%3
    n := v%5
    if v <= 0 {
      fmt.Println("-1")
    } else if m == 0 {
      for a := 0; a < v; a++ {
        fmt.Print("3")
      }
      fmt.Println("")
    } else if n == 0 {
      for a := 0; a < v; a++ {
        fmt.Print("5")
      }
      fmt.Println("")
    } else if m%5 == 0 {
      q := v-m
      fmt.Println("m :", m)
      fmt.Println("n :", n)
      fmt.Println("q :", q)
      for a := 0; a < q; a++ {
        fmt.Print("3")
      }
      for a := 0; a < m; a++ {
        fmt.Print("5")
      }
      fmt.Println("")
    } else if n%3 == 0 {
      q := v-m
      fmt.Println("m :", m)
      fmt.Println("n :", n)
      fmt.Println("q :", q)
      for a := 0; a < q; a++ {
        fmt.Print("5")
      }
      for a := 0; a < n; a++ {
        fmt.Print("3")
      }
      fmt.Println("")
    } else {
      fmt.Println("m :", m)
      fmt.Println("n :", n)
      fmt.Println("-1")
    }
  }
}
