package main

import "fmt"

var result string

func divisible(value int) string {

  if value <= 2 {
    data := "-1"
    result = ""
    return data
  } else if value%3 == 0 {
    for i:=0; i<value; i++ {
      result += "5"
    }
    data := result
    result = ""
    return data
  } else if value%5 == 0 {
    for index := 0; index < value; index++ {
      result += "3"
    }
    data := result
    result = ""
    return data
  } else {
    temp := value - 3
    result += "555"
    return result + divisible(temp)
  }
}

func chose(v int) {
  if v <= 10 {
    fmt.Println(divisible(v))
  } else {
    sum(v)
  }
}

func sum(v int) {
  temp := v - 5

  for i := temp; i > 0; i = i-5 {
    if i%3 == 0 {
      for index := 0; index < temp; index++ {
          fmt.Print("5")
      }
      ex := v - i
      for index := 0; index < ex; index++ {
        fmt.Print("3")
      }
      fmt.Println("\n")
      break
    } else {
      fmt.Println("-1")
    }
  }
}

func main() {
  var n int
  var v int

  fmt.Println("Input Value n :")
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    v = 0
    result = ""
    fmt.Println("Input Value :")
    fmt.Scan(&v)

    chose(v)
  }
}
