package main

import (
	"fmt"
	"github.com/ooyala/go-dogstatsd"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func public(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html>Hello from the public service<br/>")
	io.WriteString(w, getBackend())
	io.WriteString(w, "</html>")
}

func getBackend() string {
	backendUrl, _ := os.LookupEnv("BACKEND_URL")
	resp, err := http.Get(backendUrl)
	if err != nil {
		statsd("backend.call", "FAIL")
		return "Error contacting backend"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	statsd("backend.call", "OK")
	return string(body)
}

func statsd(title, content string) {
	statsd, err := dogstatsd.New("127.0.0.1:8125")
	defer statsd.Close()
	if err != nil {
		fmt.Println("FATAL: No statsd!")
		os.Exit(1)
	}
	statsd.Namespace = "servicechain."
	statsd.Success(title, content, nil)
}

func main() {
	http.HandleFunc("/", public)
	http.ListenAndServe(":8080", nil)
}
