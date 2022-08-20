package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	hn, err := os.Hostname()
	if err == nil {
		fmt.Fprintf(w, "hostname: %v\n", hn)
	} else {
		fmt.Fprintf(w, "%v\n", err)
	}
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", headers)

	log.Fatal(http.ListenAndServe(":1111", nil))
}
