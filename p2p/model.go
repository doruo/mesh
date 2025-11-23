package p2p

import "fmt"

// A Peer instance serves both as a server and a client. It can connect to other peers and register them.
type Peer struct {
	name  string
	host  string            `default:"localhost"`
	port  int               `default:"443"`
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

// Save peer adress in peers list.
func (p *Peer) setPeer(name string, adr string) {
	fmt.Printf("\nRegistering %s (%s)", name, adr)
	p.peers[name] = adr
}

func (p *Peer) getHello() string {
	return fmt.Sprintf("HELLO %s %s %d", p.name, p.host, p.port)
}

// Fetch peer adress in peers list.
func (p *Peer) getPeerAdr(peerName string) string {
	return p.peers[peerName]
}
