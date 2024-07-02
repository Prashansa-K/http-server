package main

import (
	"fmt"
	"net/http"
	"strings"
)

func serveEcho(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")

	if len(s) > 1 && s[2] != "" {
		bodySize := fmt.Sprintf("%d", len(s[2]))

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", bodySize)
		w.Header()["Date"] = nil
		w.Header()["Server"] = nil

		// headers should be set before WriteHeader and Write functions (unless we are sending 1xx codes)
		// otherwise default ones would be sent
		// trailers can be added afterwards too
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(s[2]))
	}
}

func serveUserAgent(w http.ResponseWriter, r *http.Request) {
	headers := r.Header

	var response string

	userAgentHeaderKey := http.CanonicalHeaderKey("User-Agent")

	if headers[userAgentHeaderKey] != nil && len(headers[userAgentHeaderKey]) > 0 {
		response = headers[userAgentHeaderKey][0]

	} else {
		response = "No user-agent header found"
	}

	w.Header()["Date"] = nil
	w.Header()["Server"] = nil

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(response)))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
