package main

import (
	"net/http"
	"regexp"
)

type Router struct {
	routes []RouteEntry
}

type RouteEntry struct {
	Method      string
	Path        *regexp.Regexp
	HandlerFunc http.HandlerFunc
}

func (re *RouteEntry) Match(r *http.Request) bool {
	if re.Method != r.Method {
		return false
	}

	if match := re.Path.FindStringSubmatch(r.URL.Path); match == nil {
		return false
	}

	return true
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Find path match
	for _, route := range rtr.routes {
		match := route.Match(r)

		if !match {
			continue
		}

		route.HandlerFunc.ServeHTTP(w, r)
		return
	}

	// No path matches
	w.Write([]byte(NOT_FOUND_RESPONSE))
}

func (rtr *Router) Route(method string, path string, handlerFunc http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")

	e := RouteEntry{
		Method:      method,
		Path:        exactPath,
		HandlerFunc: handlerFunc,
	}

	rtr.routes = append(rtr.routes, e)
}
