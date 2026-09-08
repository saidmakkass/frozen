package server

import (
	"fmt"
	"net"
)

func RunServer(addr string) net.Listener {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("Error listening on %s\n", addr))
	}
	fmt.Printf("Listening on %s\n", addr)
	for {
		con, err := ln.Accept()
		if err != nil {
			panic(fmt.Sprintf("Error during connection from %s\n", con.RemoteAddr()))
		}
		go handleConnection(con)
	}
}
