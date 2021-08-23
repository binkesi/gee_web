package gee

import (
	"log"
	"net/http"
)

type router struct {
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handler: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Add router %s - %s", method, pattern)
	key := method + "-" + pattern
	r.handler[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s", c.Path)
	}
}
