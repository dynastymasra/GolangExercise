package main

import (
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
