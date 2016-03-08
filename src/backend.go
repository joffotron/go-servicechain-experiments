package main

import (
	"io"
	"net/http"
)

func backend(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<b>Backend responding!</b>")
}

func main() {
	http.HandleFunc("/", backend)
	http.ListenAndServe(":8000", nil)
}
