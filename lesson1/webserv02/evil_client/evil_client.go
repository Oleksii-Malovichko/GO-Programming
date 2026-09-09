package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	request := "POST / HTTP/1.1\r\n" +
		"Host: localhost:5050\r\n" +
		"Content-Length: 100\r\n" +
		"\r\n" +
		"hello"

	fmt.Fprint(conn, request)

	// Закрываем соединение, не отправив обещанные 100 байт.
	conn.Close()
}