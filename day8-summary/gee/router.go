package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots   map[string]*node
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:   make(map[string]*node),
		handler: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handler[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	parms := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				parms[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				parms[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, parms
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + c.Path
		r.handler[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s", c.Path)
	}
}
