package src

import (
	"./proto"
	"fmt"
	"github.com/golang/protobuf/proto"
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
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	readPingMessage(buf[:n])
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

	_, err = conn.Write(dataToSend)
	if err != nil {
		log.Fatal("Write error", err)
	}
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

func createPingRequest() {

}

func readPingMessage(message []byte) *kademlia.PingRequest {
	newPing := &kademlia.PingRequest{}

	var err = proto.Unmarshal(message, newPing)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	//fmt.Printf("Sender: " + newPing.GetSender() + " Destination: " + newPing.GetDestination() + "\\n")
	return newPing
}

func readStoreMessage(message []byte) *kademlia.StoreRequest {
	newStore := &kademlia.StoreRequest{}

	var err = proto.Unmarshal(message, newStore)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newStore
}
