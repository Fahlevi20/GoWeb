package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website: %s\n", r.URL.Path)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
