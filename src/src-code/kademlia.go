package src

import (
	"fmt"
	kademliaProto "proto"
	"time"
)

type Kademlia struct {
	RoutingTable	RoutingTable
	Me        		Contact
	K         		int
	Alpha     		int
	HashTable 		map[string][]byte
	PingWait  		time.Duration
}

func (kademlia *Kademlia) findNode(target *Contact) {
	var contacts []Contact
	contacts,_ = kademlia.NodeLookup(target, "", false)

	if contacts[0].ID.Equals(target.ID) {
		// Found what you were looking for! :D
	} else {
		// Close enough
	}
}

func (kademlia *Kademlia) findValue(hash string) {
	var contacts []Contact
	var value string
	contacts, value = kademlia.NodeLookup(nil, hash, true)

	if value == "" {
		// No value found, want a list of contacts??? :)
	} else {
		// Return value somewhere to someone
	}
}

func (kademlia *Kademlia) Store(key string, value []byte) {
	// TODO
	kademlia.HashTable[key] = value
}

func (kademlia *Kademlia) Ping(network Network, destination string, sender Contact) {
	var found = false
	timer := time.AfterFunc(time.Second * 5, func() {
		if found == false { //Node is not found within the timer, could be dead
			//network.Node.Table.RemoveContact(sender)//todo: implement remove from bucket and make sender a contact
			//find node contact in bucket and remove it
			fmt.Printf("Could not ping node \n")
			return
		}
	})

	rpcID := network.SendPingRequest(destination, sender) //send a ping request and store the rpcID
	network.PingChannels[rpcID] = make(chan kademliaProto.PingResponse) //store a ping channel in the ping channels hash map with the rpcId as key
	response := <- network.PingChannels[rpcID] //wait for response from the ping channel

	found=true //if this code is reached a response came back and node is alive
	timer.Stop() //then timer can be stopped
	fmt.Printf("Ping RpcID: " + response.GetRpcID() + " with Response: " + response.GetResponse() + " from sender: " + response.GetSender().Address + "\n") //print the result todo: should this be printed?
	//network.Node.Table.AddContact(sender) //todo: make sender to contact, to add to front of bucket
	//find node contact in bucket and add it to front
}