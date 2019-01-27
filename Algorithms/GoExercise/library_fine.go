package main

import (
  "fmt"
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
  var a, b [3]int
  var v int
  fmt.Print("Input Date Start Day : ")
  fmt.Scan(&v)
  a[0] = v
  fmt.Print("Input Date Start Month : ")
  fmt.Scan(&v)
  a[1] = v
  fmt.Print("Input Date Start Year : ")
  fmt.Scan(&v)
  a[2] = v

  fmt.Print("Input Date End Day : ")
  fmt.Scan(&v)
  b[0] = v
  fmt.Print("Input Date End Month : ")
  fmt.Scan(&v)
  b[1] = v
  fmt.Print("Input Date End Year : ")
  fmt.Scan(&v)
  b[2] = v

  fmt.Println("Start Date :", a[0], a[1], a[2])
  fmt.Println("End Date :", b[0], b[1], b[2])

  if a[1]==b[1] && b[0]>a[0] && a[2]==b[2] {
    t := b[0] - a[0]
    h := 15 * t
    fmt.Println("Tax :", h)
  } else if a[2]==b[2] && b[1]>a[1] {
    t := b[1] - a[1]
    h := 500 * t
    fmt.Println("Tax :", h)
  } else if b[2]>a[2] {
    fmt.Println("Tax :", 10000)
  } else {
    fmt.Println("Tax :", 0)
  }
}
