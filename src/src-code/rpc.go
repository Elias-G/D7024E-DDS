package src

import (
	"github.com/golang/protobuf/proto"
	"log"
	kademliaProto "proto"
	"strconv"
	"time"
)

/*
RPC methods for Ping, Find node, Find Value and Store
*/

const wait = 2

func PingRPC(network Network, destination string, sender Contact) string {
	rpcID := NewRandomKademliaID().String()
	network.PingChannels[rpcID] = make(chan kademliaProto.PingResponse) //store a ping channel in the ping channels hash map with the rpcId as key
	SendPingRequest(destination, sender, rpcID)                         //send a ping request and store the rpcID
loop:
	for {
		select {
		case response := <-network.PingChannels[rpcID]: //wait for response from the ping channel
			log.Print(response)
			return "true"
		case <-time.After(time.Second * wait):
			break loop
		}
	}
	return "false"
}

func FindNodeRPC(network Network, destination string, targetID string, sender Contact) []Contact {
	rpcID := NewRandomKademliaID().String()
	network.FindNodeChannels[rpcID] = make(chan kademliaProto.FindNodeResponse) //store a FindNode channel in the FindNode channels hash map with the rpcId as key
	SendFindNodeRequest(destination, targetID, sender, rpcID)                   //send a FindNode request and store the rpcID
loop:
	for {
		select {
		case response := <-network.FindNodeChannels[rpcID]: //wait for response from the FindNOde channel
			return formatContactsForRead(response.Contacts) //Return the contacts
		case <-time.After(time.Second * wait):
			break loop
		}
	}
	return []Contact{}
}

func FindValueRPC(network Network, destination string, targetID string, sender Contact) ([]byte, []Contact) {
	rpcID := NewRandomKademliaID().String()
	network.FindValueChannels[rpcID] = make(chan kademliaProto.FindValueResponse) //store a FindValue channel in the FindValue channels hash map with the rpcId as key
	SendFindValueRequest(destination, targetID, sender, rpcID)                    //send a FindValue request and store the rpcID
loop:
	for {
		select {
		case response := <-network.FindValueChannels[rpcID]: //wait for response from the FindNode channel
			return response.GetValue(), formatContactsForRead(response.GetContacts())
		case <-time.After(time.Second * wait):
			break loop
		}
	}
	return nil, []Contact{}
}

func StoreRPC(network Network, destination string, sender Contact, data []byte) string {
	rpcID := NewRandomKademliaID().String()
	network.StoreChannels[rpcID] = make(chan kademliaProto.StoreResponse) //store a Store channel in the Store channels hash map with the rpcId as key
	SendStoreRequest(destination, sender, data, rpcID)                    //send a Store request and store the rpcID
loop:
	for {
		select {
		case response := <-network.StoreChannels[rpcID]: //wait for response from the Store channel
			return response.Hash //return the hash
		case <-time.After(time.Second * wait):
			break loop
		}
	}
	return "timeout. no activities under " + strconv.Itoa(wait) + " seconds"
}

/*
Functions for sending responses from Ping, Find node, Find Value and Store
*/
func sendPingResponse(rpcID string, sender Contact) []byte {
	res := &kademliaProto.PingResponse{
		RpcID:    rpcID,
		Sender:   formatContactForSend(sender),
		Response: "OK",
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	return dataToSend
}

func sendFindNodeResponse(rpcID string, sender Contact, contacts []Contact) []byte {
	res := &kademliaProto.FindNodeResponse{
		RpcID:    rpcID,
		Sender:   formatContactForSend(sender),
		Contacts: formatContactsForSend2(contacts),
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	return dataToSend
}

func sendFindValueResponse(rpcID string, sender Contact, value []byte, contacts []Contact) []byte {
	res := &kademliaProto.FindValueResponse{
		RpcID:    rpcID,
		Sender:   formatContactForSend(sender),
		Value:    value,
		Contacts: formatContactsForSend2(contacts),
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	return dataToSend
}

func sendStoreResponse(rpcID string, sender Contact, hash string) []byte {
	res := &kademliaProto.StoreResponse{
		RpcID:  rpcID,
		Sender: formatContactForSend(sender),
		Hash:   hash,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	return dataToSend
}

/*
Functions for sending requests from Ping, Find node, Find Value and Store
*/
func SendPingRequest(destination string, sender Contact, rpcID string) string {
	res := &kademliaProto.PingRequest{
		RpcID:       rpcID,
		Sender:      formatContactForSend(sender),
		Destination: destination,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, pingReqHead)
	return res.GetRpcID()
}

func SendFindNodeRequest(destination string, targetID string, sender Contact, rpcID string) string {
	res := &kademliaProto.FindNodeRequest{
		RpcID:    rpcID,
		Sender:   formatContactForSend(sender),
		TargetId: targetID,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	sendData(destination, dataToSend, findNodeReqHead)
	return res.GetRpcID()
}

func SendFindValueRequest(destination string, hash string, sender Contact, rpcID string) string {
	res := &kademliaProto.FindValueRequest{
		RpcID:  rpcID,
		Sender: formatContactForSend(sender),
		Hash:   hash,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	sendData(destination, dataToSend, findValueReqHead)
	return res.GetRpcID()
}

func SendStoreRequest(destination string, sender Contact, data []byte, rpcID string) string {
	res := &kademliaProto.StoreRequest{
		RpcID:  rpcID,
		Sender: formatContactForSend(sender),
		Value:  data,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, storeReqHead)
	return res.GetRpcID()
}

/*
Functions for reading requests from Ping, Find node, Find Value and Store
*/
func readPingRequest(message []byte) *kademliaProto.PingRequest {
	newPing := &kademliaProto.PingRequest{}

	var err = proto.Unmarshal(message, newPing)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newPing
}

func readFindNodeRequest(message []byte) *kademliaProto.FindNodeRequest {
	newFindNode := &kademliaProto.FindNodeRequest{}

	var err = proto.Unmarshal(message, newFindNode)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newFindNode
}

func readFindValueRequest(message []byte) *kademliaProto.FindValueRequest {
	newFindNode := &kademliaProto.FindValueRequest{}

	var err = proto.Unmarshal(message, newFindNode)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newFindNode
}

func readStoreRequest(message []byte) *kademliaProto.StoreRequest {
	newStore := &kademliaProto.StoreRequest{}

	var err = proto.Unmarshal(message, newStore)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newStore
}

/*
Functions for reading responses from Ping, Find node, Find Value and Store
*/
func readPingResponse(message []byte) *kademliaProto.PingResponse {
	res := &kademliaProto.PingResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readFindNodeResponse(message []byte) *kademliaProto.FindNodeResponse {
	res := &kademliaProto.FindNodeResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readFindValueResponse(message []byte) *kademliaProto.FindValueResponse {
	res := &kademliaProto.FindValueResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readStoreResponse(message []byte) *kademliaProto.StoreResponse {
	res := &kademliaProto.StoreResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func formatContactsForSend2(contacts []Contact) []*kademliaProto.Contact {
	var sendingContacts []*kademliaProto.Contact
	for _, contact := range contacts {
		sendingContacts = append(sendingContacts, &kademliaProto.Contact{NodeId: contact.ID.String(), Address: contact.Address, Distance: contact.Distance.String()})
	}
	return sendingContacts
}

func formatContactForSend(contact Contact) *kademliaProto.Contact {
	return &kademliaProto.Contact{NodeId: contact.ID.String(), Address: contact.Address, Distance: contact.Distance.String()}
}

func formatContactsForRead(contacts []*kademliaProto.Contact) []Contact {
	var readContacts []Contact
	for _, contact := range contacts {
		readContacts = append(readContacts, Contact{ID: NewKademliaID(contact.NodeId), Address: contact.Address, Distance: NewKademliaID(contact.Distance)})
	}
	return readContacts
}

func formatContactForRead(contact *kademliaProto.Contact) Contact {
	return Contact{ID: NewKademliaID(contact.NodeId), Address: contact.Address, Distance: NewKademliaID(contact.Distance)}
}
