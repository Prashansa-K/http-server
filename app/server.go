package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
)

const (
	CRLF      = "\r\n"
	ROOT_PATH = "/"
)

func Handler(conn net.Conn) {
	defer conn.Close()

	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		return
	}

	var serverResponse string
	okResponse := fmt.Sprintf("HTTP/1.1 200 OK%s%s", CRLF, CRLF)
	notFoundResponse := fmt.Sprintf("HTTP/1.1 404 Not Found%s%s", CRLF, CRLF)

	if request.URL.Path == ROOT_PATH {
		serverResponse = okResponse
	} else {
		serverResponse = notFoundResponse
	}

	_, err = conn.Write([]byte(serverResponse))
	if err != nil {
		fmt.Println("Error writing to connection: ", err.Error())
		return
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	Handler(conn)
}
