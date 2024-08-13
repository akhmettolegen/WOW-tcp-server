package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("challenge" + "\n"))

	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	nonce := string(buffer[:n])
	fmt.Printf("Server: Received nonce %s from client\n", nonce)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
