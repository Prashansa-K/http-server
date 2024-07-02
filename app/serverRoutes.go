package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func serveEcho(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")

	if len(s) > 1 && s[2] != "" {
		encoding := setEncoding(w, r)

		var response []byte

		if encoding == GZIP {
			response = compressWithGzip(s[2])
		} else {
			response = []byte(s[2])
		}

		removeDefaultHeaders(w)

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(response)))

		// headers should be set before WriteHeader and Write functions (unless we are sending 1xx codes)
		// otherwise default ones would be sent
		// trailers can be added afterwards too
		w.WriteHeader(http.StatusOK)

		w.Write(response)
	}
}

func serveUserAgent(w http.ResponseWriter, r *http.Request) {
	headers := r.Header

	var userAgentValue string
	var response []byte

	userAgentHeaderKey := http.CanonicalHeaderKey("User-Agent")

	removeDefaultHeaders(w)
	encoding := setEncoding(w, r)

	if headers[userAgentHeaderKey] != nil && len(headers[userAgentHeaderKey]) > 0 {
		userAgentValue = headers[userAgentHeaderKey][0]

	} else {
		userAgentValue = "No user-agent header found"
	}

	if encoding == GZIP {
		response = compressWithGzip(userAgentValue)
	} else {
		response = []byte(userAgentValue)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(response)))

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")

	var filename string

	if len(s) > 1 && s[2] != "" {
		filename = serverConfig.directory + s[2]
	}

	removeDefaultHeaders(w)
	encoding := setEncoding(w, r)

	fileContents, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadGateway)
		}
	}

	var response []byte

	if encoding == GZIP {
		response = compressWithGzip(string(fileContents))
	} else {
		response = fileContents
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(response)))

	w.Write(response)
}

func createFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")

	var filename string

	if len(s) > 1 && s[2] != "" {
		filename = serverConfig.directory + s[2]
	}

	removeDefaultHeaders(w)
	setEncoding(w, r)

	fmt.Println(filename)

	fileContents, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in reading request body: ", err.Error())
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error in opening file: ", err.Error())
		http.Error(w, "can't open file for writing", http.StatusBadRequest)
		return
	}

	if _, err := f.Write(fileContents); err != nil {
		fmt.Println("Error in writing to file: ", err.Error())
		http.Error(w, "can't write file", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
