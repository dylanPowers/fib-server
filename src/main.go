package main

import (
	"./fib"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(fib.FibServerHandler))
}
