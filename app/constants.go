package main

const (
	CRLF     = "\r\n"
	SPACE    = " "
	PROTOCOL = "HTTP/1.1"

	// RESPONSE CODES
	SUCCESS     = "200 OK"
	NOT_FOUND   = "404 Not Found"
	BAD_GATEWAY = "500 Bad Gateway"

	// Responses
	OK_RESPONSE          = PROTOCOL + SPACE + SUCCESS + CRLF          // "HTTP/1.1 200 OK \r\n\r\n"
	NOT_FOUND_RESPONSE   = PROTOCOL + SPACE + NOT_FOUND + CRLF + CRLF // "HTTP/1.1 404 Not Found \r\n\r\n"
	BAD_GATEWAY_RESPONSE = PROTOCOL + SPACE + BAD_GATEWAY + CRLF + CRLF
)
