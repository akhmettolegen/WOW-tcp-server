package main

import (
	"crypto/sha256"
	"fmt"
	"net"
	"strconv"
)

func solvePoW(challenge string) string {
	for i := 0; ; i++ {
		nonce := strconv.Itoa(i)
		hash := sha256.Sum256([]byte(challenge + nonce))
		fmt.Printf("Client: Trying nonce %s, hash: %x\n", nonce, hash)
		if hash[0] == 0 {
			return nonce
		}
	}
}

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

	nonce := solvePoW(challenge)
	conn.Write([]byte(nonce))

	n, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}

	fmt.Println("Server response:", string(buffer[:n]))
}
