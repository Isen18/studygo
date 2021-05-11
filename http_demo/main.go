package main

import "net/http"

func helloHandle(w http.ResponseWriter, r *http.Request) {
	text := "hello world!"
	w.Write([]byte(text))
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
