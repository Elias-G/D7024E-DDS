package src

import (
	"./proto"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
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
	buf := make([]byte, 4096)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		panic(err)
	}
	fmt.Print(" Buf: ")
	fmt.Print(buf)
	fmt.Print(" end ")
	newPing := &kademlia.PingRequest{}

	err = proto.Unmarshal(buf, newPing)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	fmt.Printf("Sender: " + newPing.GetSender() + " Destination: " + newPing.GetDestination() + "\\n")
}

func (network *Network) SendPingMessage(contact *Contact, kademliaObj Kademlia) {
	// TODO
	conn, err := net.Dial("tcp", contact.Address)
	if err != nil {
		panic(err)
	}

	ping := &kademlia.PingRequest{
		Sender:      kademliaObj.Me.Address,
		Destination: contact.Address,
	}

	dataToSend, err := proto.Marshal(ping)
	fmt.Print(" data: ")
	fmt.Print(dataToSend)
	fmt.Print(" end ")
	_, err = fmt.Print(conn, dataToSend)

	if err != nil {
		log.Fatal("Print error", err)
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
