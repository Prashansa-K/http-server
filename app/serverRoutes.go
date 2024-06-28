package main

import (
	"fmt"
	"net/http"
	"strings"
)

func serveEcho(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")
	// fmt.Println(s)

	if len(s) > 1 && s[2] != "" {
		bodySize := fmt.Sprintf("%d", len(s[2]))

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", bodySize)

		// headers should be set before WriteHeader and Write functions (unless we are sending 1xx codes)
		// otherwise default ones would be sent
		// trailers can be added afterwards too
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(s[2]))
	}
}
