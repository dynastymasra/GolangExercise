package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
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
