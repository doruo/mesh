package p2p

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// A Peer instance serves both as a server and a client. It can connect to other peers and register them.
type Peer struct {
	name string
	// Server
	adr  string `default:"localhost"`
	port int    `default:"443"`
	// Client
	peers map[string]string // name -> adress ("host:port")
}

// Returns new peer instance with no listener nor peers.
func NewPeer(name string, adr string, port int) *Peer {
	fmt.Printf("\nCreating new peer as %s, %s:%d ", name, adr, port)
	return &Peer{
		name: name,
		adr:  adr,
		port: port,
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

	buffer := make([]byte, 1024)
	length, _ := conn.Read(buffer)
	msg := string(buffer[:length])

	// Register peer connnection
	parts := strings.Split(msg, " ")
	fmt.Printf("\nHandling connection from %d", conn.RemoteAddr())

	if parts[0] == "HELLO" {

		name := parts[1]
		port := 443

		IP := getConnAdr(conn)
		adr := fmt.Sprintf("%s:%d", IP, port)

		p.registerPeer(name, adr)
		fmt.Printf("Connected to %s (%s:%d)", name, IP, port)
	} else {
		fmt.Println(msg)
	}
}

// Save peer adress in peers list.
func (p *Peer) registerPeer(name string, adr string) {
	p.peers[name] = adr
}

// CLIENT PART

// Sends message to an another peer as a client.
func (p *Peer) SendMsg(name, msg string) {

	address := p.getPeerAdr(name)
	conn := p.Connect(address)

	defer conn.Close()
	// Send message
	conn.Write([]byte(msg))
	conn.Close()
	fmt.Println(msg)
}

// Connect peer to an another peer server adress.
func (p *Peer) Connect(adr string) net.Conn {
	fmt.Printf("Connecting to %s...", adr)
	conn, err := net.Dial("tcp", adr)
	if err != nil {
		log.Fatalln("Error while connecting to server :\n", err)
	}
	fmt.Println("Success !")
	p.sendHello(conn)
	return conn
}

// Fetch peer adress in peers list.
func (p *Peer) sendHello(conn net.Conn) {
	conn.Write([]byte(p.getHello()))
}

func (p *Peer) getHello() string {
	return fmt.Sprintf("HELLO %s %d", p.name, p.port)
}

// Fetch peer adress in peers list.
func (p *Peer) getPeerAdr(peerName string) string {
	return p.peers[peerName]
}

func getConnAdr(conn net.Conn) string {
	return strings.Split(conn.RemoteAddr().String(), ":")[0]
}
