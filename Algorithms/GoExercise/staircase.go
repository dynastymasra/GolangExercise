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
  var n int

  fmt.Print("Input Max n > ")
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    for x := 0; x < n; x++ {
      if i+x > n-2 {
        fmt.Print("#")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Println("")
  }
}
