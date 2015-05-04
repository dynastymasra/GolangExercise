package main 


import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s", os)
		}
	fmt.Println()
	
	fmt.Println("When saturday?")
	today := time.Now().Weekday()
	fmt.Println("today =", today)
	fmt.Println("saturday =", time.Saturday)
	switch time.Saturday {
		case today + 0:
			fmt.Println("today")
		case today + 1:
			fmt.Println("tomorrow")
		case today + 2:
			fmt.Println("in two days")
		default:
			fmt.Println("too far away")
	}
	fmt.Println()

	fmt.Println("what time today")
	t := time.Now()
	fmt.Println("time now is", t)
	switch {
		case t.Hour() < 12:
			fmt.Println("good morning")
		case t.Hour() < 17:
			fmt.Println("good afternoon")
		default:
			fmt.Println("good evening")
	}
}

