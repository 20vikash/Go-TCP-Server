package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error connecting: ", err)
		return
	}

	defer conn.Close()

	data := []byte("Hello server")
	_, err2 := conn.Write(data)

	if err2 != nil {
		fmt.Println("Error writing: ", err2)
		return
	}
}
