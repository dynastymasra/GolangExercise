package main

import "fmt"

type IPAddr [4]byte

type Person struct {
	Name string
	Age int 
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
	fmt.Println("stringer")
	a := Person{"Arthur Dent", 42}
	b := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, b)

	fmt.Println("\nexercise stringer")
	addrs := map[string]IPAddr {
		"loopback": {127, 0, 0, 1},
		"dns": {8, 8, 8, 8},
	}
	for n, i := range addrs {
		fmt.Printf("%v: %v\n", n, i)
	}
}
