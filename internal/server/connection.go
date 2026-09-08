package server

import (
	"net"
	"fmt"
)

func handleConnection(con net.Conn) {
	fmt.Printf("Handling connection from %s\n", con.RemoteAddr())
}