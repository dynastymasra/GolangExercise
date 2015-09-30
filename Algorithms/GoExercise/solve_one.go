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
func add(a uint32, b uint32) uint32 {
  return (a+b)
}

func main()  {
  var a, b, res uint32
  fmt.Print("Input Variable > ")
  fmt.Scanf("%v %v", &a, &b)
  res = add(a, b)
  fmt.Println("Total > ", res)
}
