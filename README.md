# Word of Wisdom TCP Server

This project implements a "Word of Wisdom" TCP server that is protected against DDoS attacks using a Proof of Work (PoW) challenge-response protocol. The server sends a random quote to clients who successfully solve the PoW challenge. The project includes Docker files for both the server and the client, allowing for easy deployment.

## Overview

The Word of Wisdom TCP Server is a simple application that:

Accepts TCP connections from clients.
Issues a Proof of Work challenge to each client.
If the client solves the PoW, the server sends a random quote from a collection.
The application is Dockerized for ease of deployment.
Proof of Work
The Proof of Work (PoW) mechanism used in this project requires the client to find a nonce that, when hashed with a challenge string using the SHA-256 algorithm, produces a hash with certain properties (e.g., a leading zero byte). This ensures that only clients who expend computational effort can access the service, protecting the server against DDoS attacks.

## Usage
**Running the Server**

You can run the server either locally or using Docker.

**Locally**
```
go run ./server/main.go
```

**Using Docker**

Build the Docker image:
```
docker build -t word-of-wisdom-server -f Dockerfile.server .
```
Run the Docker container:
```
docker run -p 8080:8080 word-of-wisdom-server
```

**Running the Client**

You can also run the client locally or using Docker.

**Locally**
```
go run ./client/main.go
```
**Using Docker**

Build the Docker image
```
docker build -t word-of-wisdom-client -f Dockerfile.client .
```
Run the Docker container:
```
docker run --network="host" word-of-wisdom-client
```
## Why This PoW Algorithm?

The PoW algorithm was chosen for its balance between security and simplicity:

 - **Simplicity:** Easy to implement using cryptographic hash functions (SHA-256).

 - **Security:** Provides adequate protection against DDoS attacks by requiring computational effort to solve the challenge.

 - **Flexibility:** The difficulty can be easily adjusted by changing the hash verification condition.

 - **Proven Effectiveness:** This algorithm is widely used in applications like Bitcoin, demonstrating its effectiveness in real-world scenarios