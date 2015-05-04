package main

import ("fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("I am in Go Programming World!")
	fmt.Println("The time now is", time.Now())
	
	fmt.Println("Number random is", rand.Intn(10))
	fmt.Printf("Number next %g math between 2 and 3 \n", math.Nextafter(2, 3)) 
	fmt.Println(math.Pi)
}
