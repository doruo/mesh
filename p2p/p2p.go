package p2p

import (
	"fmt"
	"net"
	"strings"
)

type Peer struct {
	username string
	port     int
	// Server part
	listener net.Listener
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
	listener, _ := net.Listen("tcp", ":"+string(rune(p.port)))
	fmt.Println("Peer created as", p.username, " and server listening on port:", p.port)

	defer listener.Close()
	p.wait(listener)
}

func (p *Peer) wait(listener net.Listener) {
	for {
		connection, _ := listener.Accept()
		go p.handleConnection(connection)
	}
}

func (p *Peer) handleConnection(connection net.Conn) {

	buffer := make([]byte, 1024)
	length, _ := connection.Read(buffer)
	message := string(buffer[:length])

	// Register peer connnection
	parts := strings.Split(message, " ")

	if parts[0] == "HELLO" {

		name := parts[1]
		port := parts[2]

		// Fetch IP from connexion
		IP := strings.Split(connection.RemoteAddr().String(), ":")[0]
		p.register(name, IP+":"+port)

		fmt.Println("Connected to", name, "(", IP, port, ")")

	} else {
		fmt.Println(message)
	}
}

// Save peer adress in peers list.
func (p *Peer) register(peerName string, adress string) {
	p.peers[peerName] = adress
}

// - CLIENT PART -

// Sends message to an another peer.
func (p *Peer) SendMessage(peerName, message string) {

	// find peer adress
	address := p.findAdress(peerName)
	// Connect to the other peer as client
	conn, err := p.connect(address)

	if err != nil {
		fmt.Println("Error while connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send message
	conn.Write([]byte(message))
	conn.Close()
}

// Connect peer to an another peer server adress.
func (p *Peer) connect(adress string) (net.Conn, error) {

	conn, err := net.Dial("tcp", adress)
	if err != nil {
		return nil, err
	}

	// Send hello message
	hello := "HELLO " + p.username + " " + string(rune(p.port))
	conn.Write([]byte(hello))

	return conn, nil
}

// Fetch peer adress in peers list.
func (p *Peer) findAdress(peerName string) string {
	return p.peers[peerName]
}
