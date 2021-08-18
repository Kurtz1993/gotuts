package main

import (
	"fmt"
	"net/http"
)

func idxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hey there</h1>
<p>Go is fast!</p>
<p>...and simple!</p>`)
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")
}

func httpstuff() {
	http.HandleFunc("/", idxHandler)
	http.ListenAndServe(":3000", nil)
}
