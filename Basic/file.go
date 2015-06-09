package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
 )

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is sample text")
	file.WriteString("This is second test text")
	file.Close()
	stream, err := ioutil.ReadFile("sample.txt" )

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)
}