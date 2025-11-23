package p2p

import (
	"fmt"
	"log"
	"mesh/utils"
	"net"
)

// Starts client loop.
func (p *Peer) StartClient() {
	fmt.Print("Enter a peer to chat: ")
	for {
		msg := utils.ReadString()
		p.SendMsg("test1", msg)
		fmt.Println()
	}
}

// Connect peer to an another peer server adress.
func (p *Peer) Connect(adr string) net.Conn {
	fmt.Printf("Connecting to [%s] ...", adr)
	conn, err := net.Dial("tcp", adr)
	if err != nil {
		log.Fatalln("Error while connecting to peer :\n", err)
	}
	fmt.Println("\nConnection success !")
	p.sendHello(conn)
	return conn
}

// Sends message to an another peer as a client.
func (p *Peer) SendMsg(name, msg string) {

	adr := p.getPeerAdr(name)
	conn := p.Connect(adr)

	defer conn.Close()
	// Send message
	fmt.Printf("\n%s", msg)
	conn.Write([]byte(msg))
	conn.Close()
	fmt.Println(msg)
}

// Fetch peer adress in peers list.
func (p *Peer) sendHello(conn net.Conn) {
	fmt.Println("\nSending hello...")
	conn.Write([]byte(p.getHello()))
}
