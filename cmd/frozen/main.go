package main

import (
	"fmt"
	"os"
	"strconv"
	"frozen/internal/server"
)

func getAddr(args []string) (addr string) {
	const defaultPrefix = "127.0.0.1"
	lenArgs := len(args)
	if lenArgs != 2 && lenArgs != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [address] <port>\n", args[0])
		os.Exit(1)
	}
	if lenArgs == 2 {
		_, err := strconv.Atoi(args[1])
		if err != nil {
			panic(fmt.Sprintf("Invalid Port %s\n", args[1]))
		}
		addr = fmt.Sprintf("%s:%s", defaultPrefix, args[1])
	} else {
		_, err := strconv.Atoi(args[2])
		if err != nil {
			panic(fmt.Sprintf("Invalid Port %s\n", args[2]))
		}
		addr = fmt.Sprintf("%s:%s", args[1], args[2])
	}
	return
}

func main() {
	addr := getAddr(os.Args)
	server.RunServer(addr)
}