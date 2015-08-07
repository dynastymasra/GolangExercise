package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", newFunc)
	http.ListenAndServe(":2244", mux)	
}

func newFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}