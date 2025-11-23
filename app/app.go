// Package app implements a simple Mesh application interface.
package app

import (
	"fmt"
	"mesh/p2p"
	"mesh/utils"
	"net"
)

var peer *p2p.Peer

// Start main peer application instance.
func Start() {
	displayTitle()
	fmt.Println("[1] Client mode")
	fmt.Println("[2] Server mode")
	for {
		fmt.Print("\nChoose a mode : ")
		switch utils.ReadInt() {
		case 1:
			startClient()
		case 2:
			startServer()
		}
	}
}

func startServer() {
	fmt.Println("\n~ Server mode ~")
	initPeer()
	peer.StartServer()
}

func startClient() {
	fmt.Println("\n~ Client mode ~")
	initPeer()
	utils.DisplayConn(connect())
	peer.StartClient()
}

func initPeer() {
	if peer != nil {
		return
	}
	fmt.Println("Create a new peer")
	fmt.Print("Enter your name: ")
	name := utils.ReadString()
	fmt.Print("Enter your port: ")
	port := utils.ReadInt()

	peer = p2p.NewPeer(name, "localhost", port)
}

func connect() net.Conn {
	fmt.Println("\nConnect to another peer")
	fmt.Print("Enter host: ")
	host := utils.ReadString()
	fmt.Print("Enter port: ")
	port := utils.ReadInt()

	adr := fmt.Sprintf("%s:%d", host, port)
	return peer.Connect(adr)
}

func displayTitle() {
	path := "./title.txt"
	fmt.Println()
	fmt.Println(utils.ReadFile(path))
	fmt.Println()
}
