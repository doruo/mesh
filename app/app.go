package app

import (
	"fmt"
	"log"
	"mesh/p2p"
	"os"
)

var currentPeer *p2p.Peer

func Init() {
	if currentPeer != nil {
		return
	}
	displayTitle()

	fmt.Println("Create a new peer")

	fmt.Print("Enter your name: ")
	name := readString()
	fmt.Print("Enter your port: ")
	port := readInt()

	currentPeer = p2p.NewPeer(name, "localhost", port)
}

func StartServer() {
	Init()
	currentPeer.StartServer()
}

func Connect() {
	Init()
	fmt.Println("\nConnect to another peer")
	fmt.Print("Enter host: ")
	host := readString()
	fmt.Print("Enter port: ")
	port := readInt()

	adr := fmt.Sprintf("%s:%d", host, port)
	currentPeer.Connect(adr)
}

func SendMsg() {
	Connect()
	fmt.Print("Enter name: ")
	otherPeerName := readString()

	msg := fmt.Sprintf("HELLO %s", otherPeerName)
	currentPeer.SendMsg(otherPeerName, msg)
}

func readString() string {
	var input string
	fmt.Scanf("%s", &input)
	return input
}

func readInt() int {
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func displayTitle() {
	file, err := os.ReadFile("title.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
