package server

import (
	"net"
	"fmt"
	"bufio"
)

func handleConnection(con net.Conn) {
	fmt.Printf("Connected to %s\n", con.RemoteAddr())
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("received:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("read error:", err)
	}
}