package gui

import (
	"mesh/p2p"
)

func main() {
	username := "test"
	port := 8080

	peer := p2p.NewPeer(username, port)
	go peer.StartServer()

	go peer.SendMessage(username, "HELLO dorito 8080")
}
