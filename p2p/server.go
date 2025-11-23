package p2p

import (
	"fmt"
	"log"
	"mesh/utils"
	"net"
	"strings"
)

// Starts server and wait for any other peer connection.
func (p *Peer) StartServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.port))

	if err != nil {
		log.Fatalln("\nError while listening :\n", err)
	}

	fmt.Printf("\nServer listening on %d", p.port)
	defer listener.Close()
	p.wait(listener)
}

// Wait for any client connexion.
func (p *Peer) wait(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("\nError while accepting :\n", err)
		}
		go p.handleConn(conn)
	}
}

// Handle a client connexion.
func (p *Peer) handleConn(conn net.Conn) {
	msg := utils.GetConnMsg(conn)

	// Parse peer connnection
	parts := strings.Split(msg, " ")
	code := parts[0]
	name := parts[1]
	host := parts[2]
	port := parts[3]
	adr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("\nHandling connection from %s:", adr)

	// First msg
	if code == "HELLO" {
		p.setPeer(name, adr)
		fmt.Printf("\nConnected to %s (%s)", name, adr)
	} else {
		fmt.Println(msg)
	}
}
