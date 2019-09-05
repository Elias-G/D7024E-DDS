package src

import (
	"fmt"
	"net"
	"strconv"
)

type Network struct {
}

func Listen(ip string, port int) {
	// TODO
	ln, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	data := make([]byte, 512)
	n, err := conn.Read(data)
	if err != nil {
		panic(err)
	}
	s := string(data[:n])
	print(s)
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO
	conn, err := net.Dial("tcp", contact.Address)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "PING!")
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
