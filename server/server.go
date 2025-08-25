package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()
	wait(listener)
}

func wait(listener net.Listener) {
	for {
		connection, _ := listener.Accept()
		go processConnection(connection)
	}
}

func processConnection(connection net.Conn) {
	buffer := newBuffer()
	length, _ := connection.Read(buffer)
	print(buffer, length)
}

func print(buffer []byte, len int) {
	msg := string(buffer[:len])
	fmt.Println(msg)
}

func newBuffer() []byte {
	return make([]byte, 1024)
}
