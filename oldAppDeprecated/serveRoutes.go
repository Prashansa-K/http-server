package main

import (
	"fmt"
	"net/http"
)

func serveRootPath() (string, error) {
	return OK_RESPONSE, nil
}

func serveEcho(request *http.Request) (string, error) {
	fmt.Println(request.URL.Path)

	return OK_RESPONSE, nil
}
