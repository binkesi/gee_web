package main

import (
	"fmt"
	"net/http"
	"summary/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %s\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Headers[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
