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

  fmt.Print("Input Total Class :")
  fmt.Scan(&n)

  for i := 0; i < n; i++ {
    inputData()
  }
}

func inputData() {
  var n, k, v int
  var t []int

  fmt.Print("Input Total Arrive :")
  fmt.Scan(&n)
  fmt.Print("Input Total Min Arrive :")
  fmt.Scan(&k)
  for i := 0; i < n; i++ {
    fmt.Printf("Time Come [%v] :", i)
    fmt.Scan(&v)
    t = append(t, v)
  }

  printData(k, t)
}

func printData(k int, t []int) {
  late := 0
  ontime := 0
  for _, i := range t {
    if i > 0 {
      late++
    } else if i < 0 {
      ontime++
    } else {
      ontime++
    }
  }

  if ontime >= k {
    fmt.Printf("NO, Class is Started, Ontime %v, Late %v, Min %v\n", ontime, late, k)
  } else {
    fmt.Printf("YES, Class is Cancelled, Ontime %v, Late %v, Min %v\n", ontime, late, k)
  }
}
