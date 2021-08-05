package main

import (
	"fmt"
	"net/http"
	"gee"
)

func main(){
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "URL.Path = %q\n")
	})
}