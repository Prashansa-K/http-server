package main

import (
	"flag"
	"net/http"
)

var serverConfig struct {
	directory string
}

func main() {
	flag.StringVar(&serverConfig.directory, "directory", DEFAULT_DIR, "directory to access")
	flag.Parse()

	router := &Router{}

	router.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		// Auto injection of Date and Server header would cause tests to fail with "anti-cheat failed" remark
		w.Header()["Date"] = nil
		w.Header()["Server"] = nil
		w.WriteHeader(http.StatusOK)
	})

	router.Route("GET", `/echo/(?P<Message>\w*)`, serveEcho)
	router.Route("GET", `/user-agent`, serveUserAgent)
	router.Route("GET", `/files/(?P<Message>\w+)`, serveFile)

	http.ListenAndServe(":4221", router)
}
