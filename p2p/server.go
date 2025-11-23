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
		log.Fatalln("Error while listening :\n", err)
	}

	log.Printf("Server listening on %d", p.port)
	defer listener.Close()
	p.wait(listener)
}

// Wait for any client connexion.
func (p *Peer) wait(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error while accepting : ", err)
		}
		go p.handleConn(conn)
	}
}

// Handle a client connexion.
func (p *Peer) handleConn(conn net.Conn) {
	// Get adress
	adr := utils.GetConnAdr(conn)
	log.Printf("Handling connection from %s:", adr)

	// Parse peer connnection
	msg := utils.GetConnMsg(conn)
	parts := strings.Split(msg, " ")

	// First msg
	if parts[0] == "HELLO" {
		name := parts[1]
		host := parts[2]
		port := parts[3]
		adr := fmt.Sprintf("%s:%s", host, port)
		p.setPeer(name, adr)
		log.Printf("Connected to %s (%s)", name, adr)
		return
	}

	fmt.Println(msg)
}
