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
	pingRespCh chan *kademlia.PingResponse
	findNodeRespCh chan []Contact
}

var pingReqHead = []byte{0, 0, 0}
var pingResHead = []byte{0, 0, 1}
var findReqHead = []byte{0, 1, 0}
var findNodeResHead = []byte{0, 1, 1}
var findValueResHead = []byte{1, 0, 0}
var storeReqHead = []byte{1, 0, 1}
var storeResHead = []byte{1, 1, 1}

func NewNetwork(node Kademlia, pingRespCh chan *kademlia.PingResponse, findNodeRespCh chan []Contact) *Network {
	n := Network{Node: node, pingRespCh:pingRespCh, findNodeRespCh:findNodeRespCh}
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

func (network *Network) NetworkJoin(node Kademlia, rootNode Contact) {
	var table = node.Table
	//var alpha = node.Alpha
	rootNode.CalcDistance(node.Me.ID)
	table.AddContact(rootNode)

	network.NodeLookup(node.Me.ID)

	//todo: How to fill table? PING rootnode and let rootnode add it? Or FindNode on self?

	/*var closest = table.FindClosestContacts(node.Me.ID, alpha)

	for _, contact := range closest { //is this needed??
		table.AddContact(contact)
	}*/
}

func printContacts (contacts []Contact) {
	for _, contact := range contacts {
		fmt.Printf("ID: " + contact.ID.String() + " IP: " + contact.Address + "\n")
	}
}

func printContact (contact Contact) {
	fmt.Printf("ID: " + contact.ID.String() + " IP: " + contact.Address + "\n")
}

// Sends out alpha RPCs for FindNode and gets k contacts from each
func (network *Network) NodeLookup(id *KademliaID)(contacts []Contact) {
	var table = network.Node.Table
	var alpha = network.Node.Alpha
	var closest = table.FindClosestContacts(id, alpha)

	//PRINTOUTS////////////////
	fmt.Printf("CLOSEST: \n")
	printContacts(closest)
	//PRINTOUTS////////////////

	var closestSoFar = closest[0].ID
	var receivedContacts []Contact

	for i := 0; i < alpha; i++ { //todo: parallelism
		var contact = closest[i]
		network.SendFindNodeRequest(contact.Address, network.Node.Me.ID.String(), network.Node.Me) //send to one of the closest contacts: destination, target id, sender
		received := <-network.findNodeRespCh //todo: timeout
		fmt.Printf(string(len(received)))
		receivedContacts = append(receivedContacts, received...)
	}

	//PRINTOUTS////////////////
	fmt.Printf("RECIEVEDCONTACTS: \n")
	printContacts(receivedContacts)

	fmt.Printf("\n")
	//PRINTOUTS////////////

	//Add all received contacts
	for _, contact := range receivedContacts {
		network.Node.Table.AddContact(contact)
	}

	// Sort received list of contacts
	candidates := ShortList{id, receivedContacts}
	candidates.Sort()
	var shortList = candidates.Contacts

	//PRINTOUTS////////////////
	fmt.Printf("CANDIDATES: \n")
	printContacts(candidates.Contacts)
	fmt.Printf("\n")
	fmt.Printf("SHORTLIST: \n")
	printContacts(shortList)
	//PRINTOUTS////////////

	// While target ID is not yet found and recent responses are closer than the previous closest,
	// Send new find contact requests
	for !shortList[0].ID.Equals(id) && shortList[0].ID.CalcDistance(id).Less(closestSoFar.CalcDistance(id)) {
		closestSoFar = shortList[0].ID
		for i := 0; i < alpha; i++ {
			var contact = shortList[i]
			network.SendFindNodeRequest(contact.Address, id.String(), network.Node.Me)
			received := <-network.findNodeRespCh //todo: timeout on request if node is dead, as ping in kademlia.go
			fmt.Printf(string(len(received)))
			receivedContacts = append(receivedContacts, received...)
		}
	}

	//PRINTOUTS////////////////
	fmt.Printf("Returning: \n")
	printContacts(contacts)
	//PRINTOUTS////////////////

	return contacts
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
			fmt.Printf("Ping Request, Destination: " + pingRequest.GetDestination() + ", Sender: " + pingRequest.GetSender())
			sendPingResponse(pingRequest.GetSender(), network.Node.Me.Address)
		case bytes.Equal(buff, pingResHead):
			pingResponse := readPingResponse(buf[3:n])
			network.pingRespCh <- pingResponse
			//fmt.Print(pingResponse)
		//Find
		case bytes.Equal(buff, findReqHead):
			findRequest := readFindNodeRequest(buf[3:n])
			// Get NodeID as string and convert it to type KademliaID
			var targetID = NewKademliaID(findRequest.TargetId)



			//Add to routing table, if it already exists it will be moved to front of bucket by add
			var newContact = Contact{Address:findRequest.GetSender().Address, ID:NewKademliaID(findRequest.GetSender().NodeId)}
			newContact.CalcDistance(network.Node.Me.ID)

			fmt.Printf("\nADD TO BUCKET\n")
			printContact(newContact)

			network.Node.Table.AddContact(newContact)

			// List of k closest contacts to the target
			var contacts = network.Node.Table.FindClosestContacts(targetID, network.Node.K)

			fmt.Printf("CONTACTS")
			if len(contacts)>0 {
				fmt.Printf(" NOT EMPTY ")
			} else {
				fmt.Printf(" EMPTY ")
			}
			fmt.Printf("\n")
			// Send response with address of sender and list of IDs
			sendFindNodeResponse(findRequest.GetSender().Address, network.Node.Me, contacts)
		case bytes.Equal(buff, findNodeResHead):
			findNodeResponse := readFindNodeResponse(buf[3:n])
			fmt.Print(findNodeResponse)
			// TODO: Return value or list of contacts??
			var contacts = formatContactsForReading(findNodeResponse.Ids)

			fmt.Printf("CONTACTS")
			if len(contacts)>0 {
				fmt.Printf(" NOT EMPTY ")
			} else {
				fmt.Printf(" EMPTY ")
			}
			fmt.Printf("\n")

			network.findNodeRespCh <- contacts

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
		log.Printf(err.Error())
		return
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


func sendFindNodeResponse(destination string, sender Contact, ids []Contact) {
	res := &kademlia.FindNodeResponse{
		Sender: formatContactForSending(sender),
		Ids: formatContactsForSending2(ids),
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	fmt.Printf("SENDER: " + res.GetSender().Address + "\n")
	fmt.Printf("DEST: " + destination + "\n")
	for _, contact := range res.Ids {
		fmt.Printf("Address: " + contact.Address + " ID: " + contact.NodeId + "\n")
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

func (network *Network) SendPingRequest(destination string, sender string){
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

func (network *Network) SendFindNodeRequest(destination string, targetID string, sender Contact) {
	res := &kademlia.FindNodeRequest{
		Sender:		formatContactForSending(sender),
		TargetId: 	targetID,
	}
	dataToSend, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("Marshal error", err)
	}
	fmt.Printf(" SendFindNodeRequest: " + destination)
	sendData(destination, dataToSend, findReqHead)
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




func formatContactsForSending(contacts []*Contact) []*kademlia.Contact{
	var sendingContacts []*kademlia.Contact
	for _, contact := range contacts {
		sendingContacts = append(sendingContacts, &kademlia.Contact{NodeId:contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()})
	}
	return sendingContacts
}

func formatContactsForSending2(contacts []Contact) []*kademlia.Contact{
	var sendingContacts []*kademlia.Contact
	for _, contact := range contacts {
		sendingContacts = append(sendingContacts, &kademlia.Contact{NodeId:contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()})
	}
	return sendingContacts
}

func formatContactForSending(contact Contact) *kademlia.Contact{
	return &kademlia.Contact{NodeId:contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()}
}

func formatContactsForReading(contacts []*kademlia.Contact) []Contact{
	var readContacts []Contact
	for _, contact := range contacts {
		readContacts = append(readContacts, Contact{ID:NewKademliaID(contact.NodeId), Address:contact.Address, Distance:NewKademliaID(contact.Distance)})
	}
	return readContacts
}

func formatContactForReading(contact kademlia.Contact) Contact{
	return Contact{ID:NewKademliaID(contact.NodeId), Address:contact.Address, Distance:NewKademliaID(contact.Distance)}
}
