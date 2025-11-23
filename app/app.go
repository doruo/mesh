package app

import (
	"fmt"
	"log"
	"mesh/p2p"
	"os"
)

func Start() {
	displayTitle()

	username := "peerTest"
	port := 8080

	peer := p2p.NewPeer(username, port)
	peer.StartServer()

	msg := "HELLO peerTest2 8080"
	peer.SendMessage(username, msg)
}

func displayTitle() {
	file, err := os.ReadFile("title.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
