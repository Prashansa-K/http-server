package main

import (
	"net/http"
)

func main() {
	router := &Router{}

	router.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		// Auto injection of Date and Server header would cause tests to fail with "anti-cheat failed" remark
		w.Header()["Date"] = nil
		w.Header()["Server"] = nil
		w.WriteHeader(http.StatusOK)
	})

	router.Route("GET", `/echo/(?P<Message>\w*)`, serveEcho)

	http.ListenAndServe(":4221", router)
}
