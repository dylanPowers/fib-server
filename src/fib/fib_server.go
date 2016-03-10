package fib

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// FibServerHandler is the http server request handler for the Fibonacci server.
// Typically it's used as http.HandlerFunc(FibServerHandler) to create an
// http.Handler object.
func FibServerHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/fib.json" {
		http.Error(w, req.URL.Path+" was not found", http.StatusNotFound)
		return
	}
	n := req.URL.Query().Get("n")
	nint, err := strconv.Atoi(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := Fibonacci(nint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, _ := json.Marshal(res)
	io.WriteString(w, string(js))
}
