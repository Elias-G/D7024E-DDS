package src

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	kademlia "src-code/proto"
)

type Network struct {
}

var pingReqHead = []byte{0, 0, 0}
var pingResHead = []byte{0, 0, 1}
var findReqHead = []byte{0, 1, 0}
var findNodeResHead = []byte{0, 1, 1}
var findValueResHead = []byte{1, 0, 0}
var storeReqHead = []byte{1, 0, 1}
var storeResHead = []byte{1, 1, 1}

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
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	/*fmt.Print(buf)
	fmt.Print("\\n")
	fmt.Print(n)
	fmt.Print("\\n")
	fmt.Print(buf[3:n])*/

	buff := buf[:3]

	switch {
	//Ping
	case bytes.Equal(buff, pingReqHead):
		pingRequest := readPingRequest(buf[3:n])
		sendPingResponse(pingRequest.GetDestination())
	case bytes.Equal(buff, pingResHead):
		pingResponse := readPingResponse(buf[3:n])
		fmt.Print(pingResponse) //todo: what to do with the response
	//Find
	case bytes.Equal(buff, findReqHead):
		findRequest := readFindNodeRequest(buf[3:n])
		//todo: find node and respond with node id or value
		//sendFindNodeResponse(findRequest)
		fmt.Print(findRequest)
	case bytes.Equal(buff, findNodeResHead):
		findNodeResponse := readFindNodeResponse(buf[3:n])
		fmt.Print(findNodeResponse) //todo: what to do with the response
	case bytes.Equal(buff, findValueResHead):
		findValueResponse := readFindNodeResponse(buf[3:n])
		fmt.Print(findValueResponse) //todo: what to do with the response
	//Store
	case bytes.Equal(buff, storeReqHead):
		storeRequest := readStoreRequest(buf[3:n])
		//todo: hash value, store value and return hash
		sendStoreResponse(storeRequest.GetSender(), storeRequest.GetValue()) //todo: what to do with the response
	case bytes.Equal(buff, storeResHead):
		storeResponse := readStoreResponse(buf[3:n])
		fmt.Print(storeResponse) //todo: what to do with the response
	}
}

func sendData(destination string, dataToSend []byte, header []byte) {
	fmt.Print("Destination: " + destination)
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		panic(err.Error() + " : Destination " + destination)
	}
	_, err = conn.Write(append(header, dataToSend...))
	if err != nil {
		log.Fatal("Write error", err)
	}
}

func sendPingResponse(destination string) {
	res := &kademlia.PingResponse{
		Response: "OK",
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, pingResHead)
}

func sendFindNodeResponse(destination string, sender string, id []byte) {
	res := &kademlia.FindNodeResponse{
		Id:      id,
		Address: sender,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, findNodeResHead)
}

func sendFindValueResponse(destination string, value []byte) {
	res := &kademlia.FindValueResponse{
		Value: value,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, findValueResHead)

}

func sendStoreResponse(destination string, value []byte) {
	hash := ""
	//todo: hash and store here
	res := &kademlia.StoreResponse{
		Hash: hash,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, storeResHead)
}

func (network *Network) SendPingRequest(destination string, sender string) {
	res := &kademlia.PingRequest{
		Sender:      sender,
		Destination: destination,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, pingReqHead)
}

func (network *Network) SendFindContactRequest(contact *Contact, kademliaObj Kademlia) {
	/*res := &kademlia.FindNodeRequest{
		Sender:		kademliaObj.Me.Address,
		NodeId:		kademliaObj.Me.ID,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(contact.Address, dataToSend, findReqHead)*/
}

func (network *Network) SendFindDataRequest(hash string) {
	// TODO
}

func (network *Network) SendStoreRequest(contact *Contact, kademliaObj Kademlia, data []byte) {
	res := &kademlia.StoreRequest{
		Sender: kademliaObj.Me.Address,
		Value:  data,
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
