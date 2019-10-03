package src

import (
	"github.com/golang/protobuf/proto"
	"log"
	kadmeliaProto "proto"
)

/*
Functions for sending responses from Ping, Find node, Find Value and Store
*/
func sendPingResponse(rpcID string, destination string, sender Contact) {
	res := &kadmeliaProto.PingResponse{
		RpcID: rpcID,
		Sender: formatContactForSend(sender),
		Response: "OK",
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	sendData(destination, dataToSend, pingResHead)
}

func sendFindNodeResponse(rpcID string, destination string, sender Contact, contacts []Contact) {
	res := &kadmeliaProto.FindNodeResponse{
		RpcID: rpcID,
		Sender: formatContactForSend(sender),
		Contacts: formatContactsForSend2(contacts),
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	sendData(destination, dataToSend, findNodeResHead)
}

func sendFindValueResponse(rpcID string, destination string, sender Contact, value []byte) {
	res := &kadmeliaProto.FindValueResponse{
		RpcID: rpcID,
		Sender: formatContactForSend(sender),
		Value: value,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, findValueResHead)

}

func sendStoreResponse(rpcID string, destination string, sender Contact, hash string) {
	res := &kadmeliaProto.StoreResponse{
		RpcID: rpcID,
		Sender: formatContactForSend(sender),
		Hash: hash,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(destination, dataToSend, storeResHead)
}

/*
Functions for sending requests from Ping, Find node, Find Value and Store
*/
func (network *Network) SendPingRequest(destination string, sender Contact) string {
	res := &kadmeliaProto.PingRequest{
		RpcID: NewRandomKademliaID().String(),
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

func (network *Network) SendFindNodeRequest(destination string, targetID string, sender Contact) string {
	res := &kadmeliaProto.FindNodeRequest{
		RpcID: NewRandomKademliaID().String(),
		Sender:		formatContactForSend(sender),
		TargetId: 	targetID,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	sendData(destination, dataToSend, findNodeReqHead)
	return res.GetRpcID()
}

func (network *Network) SendFindValueRequest(destination string, hash string, sender Contact) string {
	res := &kadmeliaProto.FindValueRequest{
		RpcID: NewRandomKademliaID().String(),
		Sender:		formatContactForSend(sender),
		Hash: 	hash,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	sendData(destination, dataToSend, findValueReqHead)
	return res.GetRpcID()
}

func (network *Network) SendStoreRequest(contact *Contact, sender Contact, data []byte) string {
	res := &kadmeliaProto.StoreRequest{
		RpcID: NewRandomKademliaID().String(),
		Sender: formatContactForSend(sender),
		Value:  data,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	sendData(contact.Address, dataToSend, storeReqHead)
	return res.GetRpcID()
}



/*
Functions for reading requests from Ping, Find node, Find Value and Store
*/
func readPingRequest(message []byte) *kadmeliaProto.PingRequest {
	newPing := &kadmeliaProto.PingRequest{}

	var err = proto.Unmarshal(message, newPing)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newPing
}

func readFindNodeRequest(message []byte) *kadmeliaProto.FindNodeRequest {
	newFindNode := &kadmeliaProto.FindNodeRequest{}

	var err = proto.Unmarshal(message, newFindNode)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newFindNode
}

func readFindValueRequest(message []byte) *kadmeliaProto.FindValueRequest {
	newFindNode := &kadmeliaProto.FindValueRequest{}

	var err = proto.Unmarshal(message, newFindNode)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newFindNode
}

func readStoreRequest(message []byte) *kadmeliaProto.StoreRequest {
	newStore := &kadmeliaProto.StoreRequest{}

	var err = proto.Unmarshal(message, newStore)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return newStore
}




/*
Functions for reading responses from Ping, Find node, Find Value and Store
 */
func readPingResponse(message []byte) *kadmeliaProto.PingResponse {
	res := &kadmeliaProto.PingResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readFindNodeResponse(message []byte) *kadmeliaProto.FindNodeResponse {
	res := &kadmeliaProto.FindNodeResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readFindValueResponse(message []byte) *kadmeliaProto.FindValueResponse {
	res := &kadmeliaProto.FindValueResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}

func readStoreResponse(message []byte) *kadmeliaProto.StoreResponse {
	res := &kadmeliaProto.StoreResponse{}

	var err = proto.Unmarshal(message, res)

	if err != nil {
		log.Fatal("Unmarshalling error ", err)
	}
	return res
}



/*
Functions for formatting a contact or a list of contacts for request or response.
 */
func formatContactsForSend(contacts []*Contact) []*kadmeliaProto.Contact {
	var sendingContacts []*kadmeliaProto.Contact
	for _, contact := range contacts {
		sendingContacts = append(sendingContacts, &kadmeliaProto.Contact{NodeId: contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()})
	}
	return sendingContacts
}

func formatContactsForSend2(contacts []Contact) []*kadmeliaProto.Contact {
	var sendingContacts []*kadmeliaProto.Contact
	for _, contact := range contacts {
		sendingContacts = append(sendingContacts, &kadmeliaProto.Contact{NodeId: contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()})
	}
	return sendingContacts
}

func formatContactForSend(contact Contact) *kadmeliaProto.Contact {
	return &kadmeliaProto.Contact{NodeId: contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()}
}

func formatContactsForRead(contacts []*kadmeliaProto.Contact) []Contact{
	var readContacts []Contact
	for _, contact := range contacts {
		readContacts = append(readContacts, Contact{ID:NewKademliaID(contact.NodeId), Address:contact.Address, Distance:NewKademliaID(contact.Distance)})
	}
	return readContacts
}

func formatContactForRead(contact *kadmeliaProto.Contact) Contact{
	return Contact{ID:NewKademliaID(contact.NodeId), Address:contact.Address, Distance:NewKademliaID(contact.Distance)}
}

