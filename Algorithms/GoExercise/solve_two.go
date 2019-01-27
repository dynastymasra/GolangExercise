package main

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
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
