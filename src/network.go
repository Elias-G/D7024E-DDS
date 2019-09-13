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

var pingReqHead = make([]byte, 000)
var pingResHead = make([]byte, 001)
var findReqHead = make([]byte, 010)
var findNodeResHead = make([]byte, 011)
var findValueResHead = make([]byte, 100)
var storeReqHead = make([]byte, 101)
var storeResHead = make([]byte, 111)

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

func NetworkJoin(me Contact, rootNode Contact, table RoutingTable, k int) {
	table.AddContact(rootNode)
	table.FindClosestContacts(me.ID, k)
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	switch buf[:3] {
	case pingReqHead:
		pingRequest := readPingRequest(buf[4:n])
		sendPingResponse(pingRequest.GetDestination())
		fmt.Print(pingRequest)
	case findReqHead:
		findRequest := readFindNodeRequest(buf[4:n])
		//sendFindNodeResponse(findRequest.)
		fmt.Print(findRequest)
	case storeReqHead:
		storeRequest := readStoreRequest(buf[4:n])
		sendStoreResponse(storeRequest.GetSender(), storeRequest.GetValue())
		fmt.Print(storeRequest)
	case pingResHead:
		pingResponse := readPingResponse(buf[4:n])
		fmt.Print(pingResponse)
	case findNodeResHead:
		findNodeResponse := readFindNodeResponse(buf[4:n])
		fmt.Print(findNodeResponse)
	case findValueResHead:
		findValueResponse := readFindNodeResponse(buf[4:n])
		fmt.Print(findValueResponse)
	case storeResHead:
		storeResponse := readStoreResponse(buf[4:n])
		fmt.Print(storeResponse)
	}
}

func sendData(destination string, dataToSend []byte, header []byte ){
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(append(header, dataToSend))
	if err != nil {
		log.Fatal("Write error", err)
	}
}

func sendPingResponse(destination string) {
	res := &kademlia.PingResponse{
		Response:      "RESPONSE",
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, pingResHead)
}

func sendFindNodeResponse(destination string, sender string, id kademlia.FindNodeResponse_KademliaID) {
	res := &kademlia.FindNodeResponse{
		Id: 		&id,
		Address:	sender,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, findNodeResHead)
}

func sendFindValueResponse(destination string, value []byte) {
	res := &kademlia.FindValueResponse{
		Value: 		value,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, findValueResHead)

}

func sendStoreResponse(destination string, hash string) {
	res := &kademlia.StoreResponse{
		Hash: 		hash,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, storeResHead)
}

func (network *Network) SendPingRequest(contact *Contact, kademliaObj Kademlia) {
	res := &kademlia.PingRequest{
		Sender:      kademliaObj.Me.Address,
		Destination: contact.Address,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(contact.Address, dataToSend, pingReqHead)
}

func (network *Network) SendFindContactRequest(contact *Contact, kademliaObj Kademlia) {
	nodeId := &kademlia.FindNodeRequest_KademliaID{KademliaID: kademliaObj.Me.ID.}

	res := &kademlia.FindNodeRequest{
		Sender:		kademliaObj.Me.Address,
		NodeId:		nodeId,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(contact.Address, dataToSend, findReqHead)
}

func (network *Network) SendFindDataRequest(hash string) {
	// TODO
}

func (network *Network) SendStoreRequest(contact *Contact, kademliaObj Kademlia, data []byte) {
	res := &kademlia.StoreRequest{
		Sender:      kademliaObj.Me.Address,
		Value: data,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(contact.Address, dataToSend, storeReqHead)
}


func readPingRequest(message []byte) *kademlia.PingRequest {
	newPing := &kademlia.PingRequest{}

	var err = proto.Unmarshal(message, newPing)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newPing
}

func readStoreRequest(message []byte) *kademlia.StoreRequest {
	newStore := &kademlia.StoreRequest{}

	var err = proto.Unmarshal(message, newStore)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newStore
}

func readFindNodeRequest(message []byte) *kademlia.FindNodeRequest {
	newFindNode := &kademlia.FindNodeRequest{}

	var err = proto.Unmarshal(message, newFindNode)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newFindNode
}

func readPingResponse(message []byte) *kademlia.PingResponse {
	res := &kademlia.PingResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readStoreResponse(message []byte) *kademlia.StoreResponse {
	res := &kademlia.StoreResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readFindNodeResponse(message []byte) *kademlia.FindNodeResponse {
	res := &kademlia.FindNodeResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readFindValueResponse(message []byte) *kademlia.FindValueResponse {
	res := &kademlia.FindValueResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}


