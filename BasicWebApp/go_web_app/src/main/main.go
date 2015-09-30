package main

import "net/http"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", newFunc)
	http.ListenAndServe(":2244", mux)
}

func newFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}
