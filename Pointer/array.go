package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x,)
}

func main() {
	fmt.Println("array")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Println("slices")
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("slices slices")
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])
	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[4:] ==", p[4:])

	fmt.Println("making slices")
	z := make([]int, 5)
	printSlice("z", z)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	fmt.Println("nil slices")
	var h []int
	fmt.Println(h, len(h), cap(h))
	if h == nil {
		fmt.Println("nil")
	}

	fmt.Println("append")
	var  k []int
	printSlice("k", k)
	k = append(k, 0)
	printSlice("k", k)
	k = append(k, 1)
	printSlice("k", k)
	k = append(k, 2, 3, 4,)
	printSlice("k", k)

	fmt.Println("range")
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("range continued")
	meow := make([]int, 10)
	for i := range meow {
		meow[i] = 1 << uint(i)
	}
	for _, value := range meow {
		fmt.Printf("%d\n", value)
	}
	
	fmt.Println("maps")
	m = make(map[string]Vertex)
	m["test"] = Vertex{40.133, -34.34424}
	fmt.Println(m["test"])
}


