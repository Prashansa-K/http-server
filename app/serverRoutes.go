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

func serveFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")

	var filename string

	if len(s) > 1 && s[2] != "" {
		filename = serverConfig.directory + s[2]
	}

	fmt.Println(filename)

	fileContents, err := os.ReadFile(filename)
	if err != nil {
		w.Header()["Date"] = nil
		w.Header()["Server"] = nil

		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadGateway)
		}
	}

	w.Header()["Date"] = nil
	w.Header()["Server"] = nil

	w.Header().Set("Content-Type", "application/octet-stream")

	w.Write([]byte(fileContents))
}

func createFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	s := strings.Split(path, "/")

	var filename string

	if len(s) > 1 && s[2] != "" {
		filename = serverConfig.directory + s[2]
	}

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

	w.Header()["Date"] = nil
	w.Header()["Server"] = nil

	w.WriteHeader(http.StatusCreated)
}
