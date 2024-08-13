package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net"
	"time"
)

var quotes = []string{
	"Habits develop into character.",
	"Don't seek happiness–create it.",
	"Never break your promises.",
	"If you want to be happy, stop complaining",
}

func generateChallenge() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%x", rng.Int63())
}

func verifyPoW(challenge, nonce string) bool {
	hash := sha256.Sum256([]byte(challenge + nonce))
	fmt.Printf("Server: Hash for challenge %s and nonce %s: %x\n", challenge, nonce, hash)
	return hash[0] == 0 // Require the first byte to be zero
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := generateChallenge()
	conn.Write([]byte(challenge + "\n"))

	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	nonce := string(buffer[:n])
	fmt.Printf("Server: Received nonce %s from client\n", nonce)

	if verifyPoW(challenge, nonce) {
		quote := quotes[rand.Intn(len(quotes))]
		conn.Write([]byte("QUOTE:" + quote + "\n"))
	} else {
		conn.Write([]byte("INVALID POW\n"))
	}
}

func main() {
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
