package main

import (
	"net/http"
)

type route struct {
	routeName string
}

const (
	ROOT_PATH = "/"
	ECHO      = "/echo/*"
)

func ServeRequest(request *http.Request) (string, error) {
	requestPath := request.URL.Path

	// if !slices.Contains(ROUTES, route{requestPath}) {
	// 	return NOT_FOUND_RESPONSE, nil
	// }

	switch requestPath {
	case ROOT_PATH:
		return serveRootPath()
	case ECHO:
		return serveEcho(request)
	default:
		// return "", errors.New("PATH NOT REGISTERED WITH A HANDLER")
		return NOT_FOUND_RESPONSE, nil
	}
}
