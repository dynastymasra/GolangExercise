package main

import "fmt"

func main() {
  var a int
  var b int
  var n int

  fmt.Print("Input Array n > ")
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    var sum int
    fmt.Print("Input Variable > ")
    fmt.Scan(&a, &b)
    sum = a + b
    fmt.Println("Total > ",   uint(sum))
  }
}
