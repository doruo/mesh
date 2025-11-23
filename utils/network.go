package utils

import (
	"fmt"
	"net"
	"strings"
)

func DisplayConn(conn net.Conn) {
	fmt.Println("\n------------------")
	fmt.Println("~ Connexion ~")
	fmt.Println("\n------------------")
	fmt.Println(conn.LocalAddr().Network())
	fmt.Println(conn.LocalAddr().String())
	fmt.Println("\n------------------")
	fmt.Println(conn.RemoteAddr().Network())
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println("\n------------------")
	fmt.Println(GetConnMsg(conn))
	fmt.Println("\n------------------")
}

func GetConnMsg(conn net.Conn) string {
	buffer := make([]byte, 1024)
	length, _ := conn.Read(buffer)
	return string(buffer[:length])
}

func GetConnAdr(conn net.Conn) string {
	return strings.Split(conn.RemoteAddr().String(), ":")[0]
}
