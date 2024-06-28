package main

import (
	"net/http"
)

func main() {
	router := &Router{}

	router.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(OK_RESPONSE + CRLF))
	})

	router.Route("GET", `/echo/(?P<Message>\w*)`, serveEcho)

	http.ListenAndServe(":4221", router)
}
