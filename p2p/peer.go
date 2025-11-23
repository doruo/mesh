package p2p

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Peer struct {
	username string
	port     int
	// Client part
	peers map[string]string // name -> adress ("host:port")
}

// Returns new peer instance with no listener nor peers.
func NewPeer(username string, port int) *Peer {
	return &Peer{
		username: username,
		port:     port,
	}
}

// - SERVER PART -

func (p *Peer) StartServer() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.port))

	if err != nil {
		log.Fatal("Error while listening :", err)
	}

	fmt.Printf("Peer created as %s and server listening on %d", p.username, p.port)
	defer listener.Close()
	p.wait(listener)
}

func (p *Peer) wait(listener net.Listener) {
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Error while accepting: ", err)
		}
		go p.handleConnection(connection)
	}
}

func (p *Peer) handleConnection(connection net.Conn) {

	buffer := make([]byte, 1024)
	length, _ := connection.Read(buffer)
	message := string(buffer[:length])

	// Register peer connnection
	parts := strings.Split(message, " ")
	fmt.Printf("DEBUG %v", parts)

	if parts[0] == "HELLO" {

		name := parts[1]
		port := parts[2]

		// Fetch IP from connexion
		IP := getConnIPAdress(connection)
		completeAdress := fmt.Sprintf("%s:%s", IP, port)
		p.setPeer(name, completeAdress)
		fmt.Printf("Connected to %s (%s:%s)", name, IP, port)
	} else {
		fmt.Println(message)
	}
}

// Save peer adress in peers list.
func (p *Peer) setPeer(peerName string, adress string) {
	p.peers[peerName] = adress
}

// - CLIENT PART -

// Sends message to an another peer.
func (p *Peer) SendMessage(peerName, message string) {

	// Connect to corresponding peer
	address := p.getPeerAdress(peerName)
	conn := p.connect(address)

	defer conn.Close()
	// Send message
	conn.Write([]byte(message))
	conn.Close()
}

// Connect peer to an another peer server adress.
func (p *Peer) connect(adr string) net.Conn {

	conn, err := net.Dial("tcp", adr)
	if err != nil {
		log.Fatal("Error while connecting to server :", err)
	}

	p.sendHelloMsg(conn)
	return conn
}

// Fetch peer adress in peers list.
func (p *Peer) sendHelloMsg(conn net.Conn) {
	conn.Write([]byte(p.getHelloMsg()))
}

func (p *Peer) getHelloMsg() string {
	return fmt.Sprintf("HELLO %s %d", p.username, p.port)
}

// Fetch peer adress in peers list.
func (p *Peer) getPeerAdress(peerName string) string {
	return p.peers[peerName]
}

func getConnIPAdress(conn net.Conn) string {
	return strings.Split(conn.RemoteAddr().String(), ":")[0]
}
