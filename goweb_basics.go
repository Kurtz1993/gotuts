package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello there")
}

func web() {
	http.HandleFunc("/", index_handler)

	http.ListenAndServe(":8080", nil)
}
