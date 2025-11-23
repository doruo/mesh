package p2p

import (
	"fmt"
	"log"
	"mesh/utils"
	"net"
	"strings"
)

// A Peer instance serves both as a server and a client. It can connect to other peers and register them.
type Peer struct {
	name string
	// Server
	host string `default:"localhost"`
	port int    `default:"443"`
	// Client
	peers map[string]string // name -> adress ("host:port")
}

// Returns new peer instance with no listener nor peers.
func NewPeer(name string, host string, port int) *Peer {
	fmt.Printf("\nCreating new peer as %s, %s:%d ", name, host, port)
	return &Peer{
		name:  name,
		host:  host,
		port:  port,
		peers: make(map[string]string, 20),
	}
}

// SERVER

// Starts server and wait for any other peer connection.
func (p *Peer) StartServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.port))

	if err != nil {
		log.Fatalln("Error while listening :\n", err)
	}

	fmt.Printf("\nServer listening on %d", p.port)
	defer listener.Close()
	p.wait(listener)
}

func (p *Peer) wait(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error while accepting :\n", err)
		}
		go p.handleConn(conn)
	}
}

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
		p.registerPeer(name, adr)
		fmt.Printf("Connected to %s (%s)", name, adr)
	} else {
		fmt.Println(msg)
	}
}

// Save peer adress in peers list.
func (p *Peer) registerPeer(name string, adr string) {
	fmt.Printf("\nRegistering %s (%s)", name, adr)
	p.peers[name] = adr
}

// CLIENT PART

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

// Connect peer to an another peer server adress.
func (p *Peer) Connect(adr string) net.Conn {
	fmt.Printf("Connecting to %s...", adr)
	conn, err := net.Dial("tcp", adr)
	if err != nil {
		log.Fatalln("Error while connecting to peer :\n", err)
	}
	fmt.Println("\nConnexion succeded !")
	p.sendHello(conn)
	return conn
}

// Fetch peer adress in peers list.
func (p *Peer) sendHello(conn net.Conn) {
	fmt.Println("\nSending hello...")
	conn.Write([]byte(p.getHello()))
}

func (p *Peer) getHello() string {
	return fmt.Sprintf("HELLO %s %s %d", p.name, p.host, p.port)
}

// Fetch peer adress in peers list.
func (p *Peer) getPeerAdr(peerName string) string {
	return p.peers[peerName]
}
