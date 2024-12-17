package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil { //If the error is EOF, that means the client closed the connection
			fmt.Println("Error reading: ", err)
			return
		}

		fmt.Println("Recieved: ", buffer[:n])
	}
}
