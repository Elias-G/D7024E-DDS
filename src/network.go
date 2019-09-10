package src

import (
	"fmt"
	"log"
	"net"
)

type Network struct {
}

func Listen(address string) {
	// TODO
	ln, err := net.Listen("tcp", address)
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
	_, err = fmt.Fprintf(conn, "PING!")

	if err != nil {
		log.Fatal("FPrintf error", err)
	}

	//Example of marshaling and unmarshaling
	/*ping := &kademlia.PingRequest{
		Sender: "10.0.0.3", //my address??
		Destination: contact.Address,
	}
	data, err := proto.Marshal(ping)

	if err != nil{
		log.Fatal("Marshaling error", err)
	}

	newPing := &kademlia.PingRequest{}

	err = proto.Unmarshal(data, newPing)

	if err != nil{
		log.Fatal("Unmarshaling error", err)
	}

	newPing.GetDestination()
	newPing.GetSender()*/
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
