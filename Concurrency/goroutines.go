package main

import (
	"fmt"
	// "time"
	"runtime"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
func say(s string) {
	for i := 0; i < 10; i++ {
		// time.Sleep(100 * time.Millisecond)
		runtime.GOMAXPROCS(runtime.NumCPU())
		runtime.Gosched()
		fmt.Println(s, runtime.NumGoroutine())
	}
}

func main() {
	go say("world")
	go say("golang")
	say("hello")
	fmt.Println("Number of CPU Cores:", runtime.NumCPU())
}
