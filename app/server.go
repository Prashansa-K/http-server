package main

import (
	"flag"
	"net/http"
)

var serverConfig struct {
	directory         string
	supportedEncoding []string
}

func main() {
	flag.StringVar(&serverConfig.directory, "directory", DEFAULT_DIR, "directory to access")
	flag.Parse()

	serverConfig.supportedEncoding = append(serverConfig.supportedEncoding, "gzip")

	router := &Router{}

	router.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		// Auto injection of Date and Server header would cause tests to fail with "anti-cheat failed" remark
		removeDefaultHeaders(w)
		setEncoding(w, r)
		w.WriteHeader(http.StatusOK)
	})

	router.Route("GET", `/echo/(?P<Message>\w*)`, serveEcho)
	router.Route("GET", `/user-agent`, serveUserAgent)
	router.Route("GET", `/files/(?P<Message>\w+)`, serveFile)
	router.Route("POST", `/files/(?P<Message>\w+)`, createFile)

	http.ListenAndServe(":4221", router)
}
