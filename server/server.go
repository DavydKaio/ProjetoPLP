package main

import (
	"bufio"
	"fmt"
	"log"
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client Left")
		conn.Close()
		return
	}

	log.Println("Cliente message: ", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}
