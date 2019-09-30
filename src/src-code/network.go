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
	Node Kademlia
	findNodeRespCh chan [] string
}

var pingReqHead = []byte{0, 0, 0}
var pingResHead = []byte{0, 0, 1}
var findReqHead = []byte{0, 1, 0}
var findNodeResHead = []byte{0, 1, 1}
var findValueResHead = []byte{1, 0, 0}
var storeReqHead = []byte{1, 0, 1}
var storeResHead = []byte{1, 1, 1}

func NewNetwork(node Kademlia) *Network {
	n := Network{Node: node}
	return &n
}

func (network *Network) Listen(address string) {
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
		go network.handleConnection(conn)
	}
}

func NetworkJoin(node Kademlia, rootNode Contact) {
	var table = node.Table
	var alpha = node.Alpha
	table.AddContact(rootNode)
	table.FindClosestContacts(node.Me.ID, alpha)
}

// Sends out alpha RPCs for FindNode and gets k contacts from each
func (network *Network) NodeLookup(id *KademliaID)(contacts []Contact) {
	var table = network.Node.Table
	var alpha = network.Node.Alpha
	var closest = table.FindClosestContacts(id, alpha)
	var receivedContacts []string
	var receivedIDs []KademliaID
	for i := 0; i < alpha; i++ {
		var contact = closest[i]
		network.SendFindContactRequest(&contact, network.Node, id)
		receivedContacts = append(receivedContacts, <- network.findNodeRespCh...)
	}
	receivedIDs = stringToKademliaID(receivedContacts)
	// TODO: Sort received IDs, is first contact target ID?
	if receivedIDs[0] != *id {
		// TODO: Send new find contact requests
	}
	return contacts
}

func stringToKademliaID(strings []string)(ids []KademliaID) {
	for i, str := range strings {
		ids[i] = *NewKademliaID(str)
	}
	return ids
}

func (network *Network) handleConnection(conn net.Conn) {
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	buff := buf[:3]

	switch {
		//Ping
		case bytes.Equal(buff, pingReqHead):
			pingRequest := readPingRequest(buf[3:n])
			sendPingResponse(pingRequest.GetDestination(), network.Node.Me.Address)
		case bytes.Equal(buff, pingResHead):
			pingResponse := readPingResponse(buf[3:n])
			fmt.Print(pingResponse) //todo: what to do with the response
		//Find
		case bytes.Equal(buff, findReqHead):
			findRequest := readFindNodeRequest(buf[3:n])
			// Get NodeID as string and convert it to type KademliaID
			var id = NewKademliaID(findRequest.NodeId)
			// List of k closest contacts to the target
			var contacts = network.Node.Table.FindClosestContacts(id, network.Node.K)
			// List of IDs to the k closest contacts
			var ids []string
			for i := 0; i < network.Node.K; i++ {
				ids[i] = contacts[i].ID.String()
			}
			// Send response with address of sender and list of IDs
			sendFindNodeResponse(findRequest.GetSender(), network.Node.Me.Address, ids)
			fmt.Print(findRequest)
		case bytes.Equal(buff, findNodeResHead):
			findNodeResponse := readFindNodeResponse(buf[3:n])
			// TODO: Return value or list of contacts??
			network.findNodeRespCh <- findNodeResponse.Ids

			fmt.Print(findNodeResponse)
		case bytes.Equal(buff, findValueResHead):
			findValueResponse := readFindNodeResponse(buf[3:n])
			fmt.Print(findValueResponse) //todo: what to do with the response
		//Store
		case bytes.Equal(buff, storeReqHead):
			storeRequest := readStoreRequest(buf[3:n])
			// Hash value and store (key, value) pair in hashtable
			key := HashValue(storeRequest.GetValue())
			network.Node.Store(key, storeRequest.GetValue())
			// Return hash
			sendStoreResponse(storeRequest.GetSender(), network.Node.Me.Address, key)
		case bytes.Equal(buff, storeResHead):
		storeResponse := readStoreResponse(buf[3:n])
		fmt.Print(storeResponse) //todo: what to do with the response
	}
}

func sendData(destination string, dataToSend []byte, header []byte) {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(append(header, dataToSend...))
	if err != nil {
		log.Fatal("Write error", err)
	}
}

func sendPingResponse(destination string, sender string) {
	res := &kademlia.PingResponse{
		Sender: sender,
		Response: "OK",
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, pingResHead)
}

func sendFindNodeResponse(destination string, sender string, ids []string) {
	res := &kademlia.FindNodeResponse{
		Sender: sender,
		Ids: ids,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, findNodeResHead)
}

func sendFindValueResponse(destination string, sender string, value []byte) {
	res := &kademlia.FindValueResponse{
		Sender: sender,
		Value: value,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, findValueResHead)

}

func sendStoreResponse(destination string, sender string, value string) {
	hash := value

	res := &kademlia.StoreResponse{
		Sender: sender,
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

// TODO: Send along target ID
func (network *Network) SendFindContactRequest(contact *Contact, kademliaObj Kademlia, targetID *KademliaID) {
	res := &kademlia.FindNodeRequest{
		Sender:		kademliaObj.Me.Address,
		NodeId:		kademliaObj.Me.ID.String(),
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
