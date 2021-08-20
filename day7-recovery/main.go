package main

import (
	"gee_recovery/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello gee\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"sungn"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
