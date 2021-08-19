package main

import (
	"net/http"
	"text/template"
)

type NewsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, req *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "Some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", newsAggHandler)

	http.ListenAndServe(":5000", nil)
}
