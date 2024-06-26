package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
)

func Handler(conn net.Conn) {
	defer conn.Close()

	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		return
	}

	serverResponse, err := ServeRequest(request)
	if err != nil {
		fmt.Println("Error serving request: ", err.Error())
		serverResponse = BAD_GATEWAY_RESPONSE
	}

	_, err = conn.Write([]byte(serverResponse))
	if err != nil {
		fmt.Println("Error writing to connection: ", err.Error())
		return
	}
}

func main() {
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
