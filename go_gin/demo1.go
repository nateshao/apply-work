package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%s hello world", r.Method)
	})
	_ = http.ListenAndServe(":8080", nil)
}
