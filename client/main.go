package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}

	challenge := string(buffer[:n])
	fmt.Printf("Client: Received challenge %s\n", challenge)
	conn.Write([]byte("nonce"))
}
