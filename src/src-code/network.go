package src

import (
	"bytes"
	"fmt"
	"log"
	"net"
	kademliaProto "proto"
	"strconv"
)

type Network struct {
	Node Kademlia
	//Channels for responses
	PingChannels map[string]chan kademliaProto.PingResponse
	FindNodeChannels map[string]chan kademliaProto.FindNodeResponse
	FindValueChannels map[string]chan kademliaProto.FindValueResponse
	StoreChannels map[string]chan kademliaProto.StoreResponse
}

//Headers for the switch in handleConnection
var pingReqHead = []byte{0, 0, 0}
var pingResHead = []byte{0, 0, 1}
var findNodeReqHead = []byte{0, 1, 0}
var findValueReqHead = []byte{0, 1, 1}
var findNodeResHead = []byte{1, 0, 1}
var findValueResHead = []byte{1, 1, 0}
var storeReqHead = []byte{1, 0, 0}
var storeResHead = []byte{1, 1, 1}

func NewNetwork(node Kademlia) *Network {
	n := Network{
		Node: node,
		PingChannels: make(map[string]chan kademliaProto.PingResponse),
		FindNodeChannels: make(map[string]chan kademliaProto.FindNodeResponse),
		FindValueChannels: make(map[string]chan kademliaProto.FindValueResponse),
		StoreChannels: make(map[string]chan kademliaProto.StoreResponse),
	}
	return &n
}

func (network *Network) Listen(address string) {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)

	fmt.Print(udpAddr)

	if err != nil {
		log.Printf(err.Error())
		return
	}

	serverConn, err := net.ListenUDP("udp", udpAddr)// &net.UDPAddr{IP:[]byte{0,0,0,0},Port:5000,Zone:""})
	if err != nil {
		log.Printf(err.Error())
		return
	}
	defer serverConn.Close()
	for {
		network.handleConnection(*serverConn) //pass connection to switch
	}
}

func sendData(destination string, dataToSend []byte, header []byte) {
	/*fmt.Printf("DataToSend: ")
	fmt.Print(dataToSend)
	fmt.Printf("\n")*/
	fmt.Printf("Destination, datatosend: " + destination + "\n")
	udpAddr, err := net.ResolveUDPAddr("udp4", destination)

	if err != nil {
		log.Printf(err.Error())
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	defer conn.Close()

	if err != nil {
		log.Printf(err.Error())
		return
	}
	_, err = conn.Write(append(header, dataToSend...))
	if err != nil {
		log.Fatal("Write error", err)
	}
}

func (network *Network) NetworkJoin(node Kademlia, rootNode Contact) {
	//var alpha = node.Alpha
	rootNode.CalcDistance(node.Me.ID)
	network.addContact(rootNode)

	//todo: Iterative find here
	shortlist := network.Node.findNode(*network, node.Me.ID.String())
	for _, contact := range shortlist {
		if contact.Address != node.Me.Address {
			network.addContact(contact)
		}
	}
	fmt.Print("This is length of shortlist: " + strconv.Itoa(len(shortlist)) + "\n")
}

func (network *Network) handleConnection(conn net.UDPConn) { //todo: this switch should contain as little code as possible, try to move functionality/logic to help functions
	message := make([]byte, 1024) //Buffer to store message received in
	n, _ , err := conn.ReadFromUDP(message) //read incoming messages
	if err != nil { //Error handling
		log.Fatal(err)
	}
	header := message[:3] //parse the header
	fmt.Print(header)
	fmt.Print("\n")
	/*fmt.Printf("handleConnection, message: ")
	fmt.Print(message)
	fmt.Printf("\n\n")*/
	switch {
		//Ping
		case bytes.Equal(header, pingReqHead):
			pingRequest := readPingRequest(message[3:n]) //read request
			fmt.Printf("Ping Request ID: " + pingRequest.GetRpcID() + " from sender: " + pingRequest.GetSender().Address + " to: " +  pingRequest.GetDestination() + "\n") //print the result todo: should this be printed?
			data := sendPingResponse(pingRequest.RpcID, network.Node.Me) //send response with rpcID from request //todo: functionality to own function
			sendData(pingRequest.GetSender().Address, data, pingResHead)
    
    
		case bytes.Equal(header, pingResHead):
			pingResponse := readPingResponse(message[3:n]) //read response
			network.PingChannels[pingResponse.RpcID]  <- *pingResponse //Get the channel with the correct rpcID from the PingChannels Hashmap in network and send the response back to that channel

		//Find Node
		case bytes.Equal(header, findNodeReqHead):
			findNodeRequest := readFindNodeRequest(message[3:n])
			network.updateRoutingTableWithoutMe(formatContactForRead(findNodeRequest.GetSender()))
			contacts := network.FindNode(findNodeRequest)
			data := sendFindNodeResponse(findNodeRequest.RpcID, network.Node.Me, contacts)
			sendData(findNodeRequest.GetSender().Address, data, findNodeResHead)
    
    
		case bytes.Equal(header, findNodeResHead):
			fmt.Printf("Find node response header switch\n")
			findNodeResponse := readFindNodeResponse(message[3:n])
			//fmt.Printf("Findnode response Request ID: " + findNodeResponse.GetRpcID() + " from sender: " + findNodeResponse.GetSender().Address + " With contacts: " + printContacts(formatContactsForRead(findNodeResponse.GetContacts())) + "\n")
			network.updateRoutingTableWithoutMe(formatContactForRead(findNodeResponse.GetSender()))
			fmt.Printf("Before response find node\n")
			network.FindNodeChannels[findNodeResponse.RpcID]  <- *findNodeResponse
			fmt.Printf("After find node response\n")


		//Find Value
		case bytes.Equal(header, findValueReqHead):
			findValueReq := readFindValueRequest(message[3:n])
			network.updateRoutingTableWithoutMe(formatContactForRead(findValueReq.GetSender()))
			value, contacts := network.FindValue(findValueReq)
			data := sendFindValueResponse(findValueReq.GetRpcID(), network.Node.Me, value, contacts)
			sendData(findValueReq.GetSender().Address, data, findValueResHead)
    
    
		case bytes.Equal(header, findValueResHead):
			findValueResponse := readFindValueResponse(message[3:n])
			network.updateRoutingTableWithoutMe(formatContactForRead(findValueResponse.GetSender()))
			network.FindValueChannels[findValueResponse.RpcID]  <- *findValueResponse


		//Store
		case bytes.Equal(header, storeReqHead):
			storeRequest := readStoreRequest(message[3:n])
      network.updateRoutingTableWithoutMe(formatContactForRead(storeRequest.GetSender()))
			hash := network.Node.Store(storeRequest.GetValue())
			data := sendStoreResponse(storeRequest.RpcID, network.Node.Me, hash)
			sendData(storeRequest.GetSender().Address, data, storeResHead)

    
		case bytes.Equal(header, storeResHead):
			fmt.Printf("Store response header switch\n")
			storeResponse := readStoreResponse(message[3:n])
			network.updateRoutingTableWithoutMe(formatContactForRead(storeResponse.GetSender()))
			fmt.Printf("Before response store, looking for " + storeResponse.RpcID + "\n")
			testChannels(network.StoreChannels, storeResponse.RpcID)
			network.StoreChannels[storeResponse.RpcID]  <- *storeResponse
			fmt.Printf("After store response\n")
	}
}

func testChannels(hashmap map[string]chan kademliaProto.StoreResponse, rpcId string) {
	if hashmap[rpcId] != nil {
		fmt.Printf("Channel not null\n")
	}   else {
		fmt.Printf("Channel is null\n")
	}
}

func (network *Network) FindNode(findNodeRequest *kademliaProto.FindNodeRequest) []Contact {
	// Get NodeID as string and convert it to type KademliaID
	var targetID = NewKademliaID(findNodeRequest.TargetId)
	//Add to routing table, if it already exists it will be moved to front of bucket by add
	var newContact = Contact{Address:findNodeRequest.GetSender().Address, ID:NewKademliaID(findNodeRequest.GetSender().NodeId)}
	newContact.CalcDistance(network.Node.Me.ID)
	network.addContact(newContact)
	// List of k closest contacts to the target
	contacts := network.Node.RoutingTable.FindClosestContacts(targetID, network.Node.K)
	contacts = dontAddRequester(contacts, formatContactForRead(findNodeRequest.GetSender())) //Don't add requester to response
	fmt.Printf("FindNode network, contacts back: " + printContacts(contacts) + "\n")
	return contacts
}

func (network *Network) FindValue(findValueRequest *kademliaProto.FindValueRequest)(value []byte, contacts []Contact) {
	// Get NodeID as string and convert it to type KademliaID
	var hash = NewKademliaID(findValueRequest.Hash)
	//Add to routing table, if it already exists it will be moved to front of bucket by add
	var newContact = Contact{Address:findValueRequest.GetSender().Address, ID:NewKademliaID(findValueRequest.GetSender().NodeId)}
	newContact.CalcDistance(network.Node.Me.ID)
	network.addContact(newContact)

	if val, ok := network.Node.HashTable[findValueRequest.Hash]; ok {
		value = val
	} else {
		// List of k closest contacts to the target
		contacts = network.Node.RoutingTable.FindClosestContacts(hash, network.Node.K)
		contacts = dontAddRequester(contacts, formatContactForRead(findValueRequest.GetSender())) //Don't add requester to response
	}
	return value, contacts
}

func dontAddRequester(contacts []Contact, sender Contact) []Contact {
	for i, contact := range contacts {
		if contact.ID == sender.ID{
			return append(contacts[:i], contacts[i+1])
		}
	}
	return contacts
}

func (network *Network) addContact(contact Contact){
	if !network.Node.Me.ID.Equals(contact.ID) { //don't add yourself
		network.Node.RoutingTable.AddContact(contact)
	}
}

func (network *Network) updateRoutingTableWithoutMe(contact Contact) {
	if !network.Node.Me.ID.Equals(contact.ID) { //don't add yourself
		network.Node.RoutingTable.UpdateRoutingTable(contact)
	}
}

func GetIpAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("interface error", err)
	}

	for _, i := range ifaces {
		if i.Name == "eth0" {
			addrs, err := i.Addrs()
			if err != nil {
				log.Fatal("interface error", err)
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				return ip.String()
			}
		}
	}
	return ""
}

func CreateNode(port int, ip string, id *KademliaID) Contact {
	address := ip + ":" + strconv.Itoa(port)
	me := NewContact(id, address)
	return me
}

func printContacts(contacts []Contact) string {
	s := ""
	for _, contact := range contacts {
		s += contact.Address + ", "
	}
	return s
}