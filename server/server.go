package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	ln, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error Listening: " + err.Error())
		os.Exit(1)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error Conecting: " + err.Error())
			return
		}

		fmt.Println("Client Connected")

		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")
	}
}
