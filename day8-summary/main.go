package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL path is %s\n", req.URL.Path)
	})
	http.ListenAndServe(":9999", nil)
}
