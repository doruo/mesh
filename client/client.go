package main

import (
	"fmt"
	"net"
	"os"
)

var server string = "127.0.0.1:8080"

func main() {

	// Verify arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run client.go \"message\"")
		fmt.Println("Exemple: go run client.go \"Hello World\"")
		return
	}

	// Connect to server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error while connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send message to server
	msg := []byte(os.Args[1])
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("Error while sending msg to server:", err)
		return
	}
}
