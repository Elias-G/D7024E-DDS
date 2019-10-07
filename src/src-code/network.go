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
	ln, err := net.Listen("tcp", address) //Listener
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept() //accept incoming connection
		if err != nil {
			panic(err)
		}
		go network.handleConnection(conn) //pass connection to switch
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

func (network *Network) NetworkJoin(node Kademlia, rootNode Contact) {
	var table = node.RoutingTable
	//var alpha = node.Alpha
	rootNode.CalcDistance(node.Me.ID)
	table.AddContact(rootNode)

	//todo: Iterative find here
	shortlist, _ := network.Node.NodeLookup(*network, &node.Me , node.Me.ID.String(), false)
	fmt.Print("This is length of shortlist: " + string(len(shortlist)) + "\n")
}

func (network *Network) handleConnection(conn net.Conn) { //todo: this switch should contain as little code as possible, try to move functionality/logic to help functions
	message := make([]byte, 512) //Buffer to store message received in
	n, err := conn.Read(message) //read incoming messages
	if err != nil { //Error handling
		log.Fatal(err)
	}
	header := message[:3] //parse the header
	switch {
		//Ping
		case bytes.Equal(header, pingReqHead):
			pingRequest := readPingRequest(message[3:n]) //read request
			fmt.Printf("Ping Request ID: " + pingRequest.GetRpcID() + " from sender: " + pingRequest.GetSender().Address + " to: " +  pingRequest.GetDestination() + "\n") //print the result todo: should this be printed?
			sendPingResponse(pingRequest.RpcID, pingRequest.GetSender().Address, network.Node.Me) //send response with rpcID from request //todo: functionality to own function


		case bytes.Equal(header, pingResHead):
			pingResponse := readPingResponse(message[3:n]) //read response
			network.PingChannels[pingResponse.RpcID]  <- *pingResponse //Get the channel with the correct rpcID from the PingChannels Hashmap in network and send the response back to that channel


		//Find Node
		case bytes.Equal(header, findNodeReqHead): //todo: move code to own function
			findNodeRequest := readFindNodeRequest(message[3:n])
			success := network.Node.RoutingTable.UpdateRoutingTable(formatContactForRead(findNodeRequest.GetSender()))
			if(!success){
				
			}
			// Get NodeID as string and convert it to type KademliaID
			var targetID = NewKademliaID(findNodeRequest.TargetId)
			//Add to routing table, if it already exists it will be moved to front of bucket by add
			var newContact = Contact{Address:findNodeRequest.GetSender().Address, ID:NewKademliaID(findNodeRequest.GetSender().NodeId)}
			newContact.CalcDistance(network.Node.Me.ID)
			// List of k closest contacts to the target
			var contacts = network.Node.RoutingTable.FindClosestContacts(targetID, network.Node.K)
			// Send response with address of sender and list of IDs
			sendFindNodeResponse(findNodeRequest.RpcID, findNodeRequest.GetSender().Address, network.Node.Me, contacts)


		case bytes.Equal(header, findNodeResHead):
			findNodeResponse := readFindNodeResponse(message[3:n])
			fmt.Printf("Findnode response Request ID: " + findNodeResponse.GetRpcID() + " from sender: " + findNodeResponse.GetSender().Address + "\n")
			network.Node.RoutingTable.UpdateRoutingTable(formatContactForRead(findNodeResponse.GetSender()))
			network.FindNodeChannels[findNodeResponse.RpcID]  <- *findNodeResponse
		//Find Value
		
		
		case bytes.Equal(header, findValueReqHead): //todo: functionality to own function when implemented
			findValueReq := readFindValueRequest(message[3:n])
			network.Node.RoutingTable.UpdateRoutingTable(formatContactForRead(findValueReq.GetSender()))
			fmt.Print(findValueReq) //todo: remove, only here such that var is used right now
		
		
		case bytes.Equal(header, findValueResHead):
			findValueResponse := readFindValueResponse(message[3:n])
			network.Node.RoutingTable.UpdateRoutingTable(formatContactForRead(findValueResponse.GetSender()))
			network.FindValueChannels[findValueResponse.RpcID]  <- *findValueResponse
		
		
			//Store
		case bytes.Equal(header, storeReqHead): //todo: move code to own function
			storeRequest := readStoreRequest(message[3:n])
			network.Node.RoutingTable.UpdateRoutingTable(formatContactForRead(storeRequest.GetSender()))
			// Hash value and store (key, value) pair in hashtable
			key := HashValue(storeRequest.GetValue())
			network.Node.Store(key, storeRequest.GetValue())
			// Return hash
			sendStoreResponse(storeRequest.RpcID, storeRequest.GetSender().Address, network.Node.Me, key)


		case bytes.Equal(header, storeResHead):
			storeResponse := readStoreResponse(message[3:n])
			network.Node.RoutingTable.UpdateRoutingTable(formatContactForRead(storeResponse.GetSender()))
			network.StoreChannels[storeResponse.RpcID]  <- *storeResponse
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