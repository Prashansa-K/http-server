package main

import (
	"fmt"
	"net/http"
	"strings"
)

func serveEcho(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")
	fmt.Println(s)

	if len(s) > 1 && s[2] != "" {
		bodySize := len(s[2])
		contentType := "Content-Type: text/plain\r\n"
		contentLen := fmt.Sprintf("Content-Length: %d%s", bodySize, CRLF)
		w.Write([]byte(
			OK_RESPONSE +
				contentType +
				contentLen +
				s[2],
		))
	}

	// w.Write([]byte(OK_RESPONSE))
}
