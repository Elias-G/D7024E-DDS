package src

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Kademlia struct {
	RoutingTable RoutingTable
	Me           Contact
	K            int
	Alpha        int
	HashTable    map[string][]byte
	PingWait     time.Duration
}

/*
CLI commands
*/
func (kademlia *Kademlia) PutCommand(network Network, value []byte) {
	hash := HashValue(value)
	nodes := kademlia.findNode(network, hash)

	if network.Node.Me.ID.CalcDistance(NewKademliaID(hash)).Less(nodes[len(nodes)-1].ID.CalcDistance(NewKademliaID(hash))) { //If this node is closer than the last one in the returned list this node should replace that one
		nodes[len(nodes)-1] = network.Node.Me
	}

	for _, node := range nodes {
		StoreRPC(network, node.Address, kademlia.Me, value)
	}
	fmt.Printf(hash + "\n")
}

func (kademlia *Kademlia) GetCommand(network Network, hash string) {
	value := kademlia.findValue(network, hash)
	fmt.Printf(value + "\n")
}

func (kademlia *Kademlia) ExitCommand() {
	os.Exit(1)
}

/*
RPCs
*/
func (kademlia *Kademlia) findNode(network Network, target string) []Contact {
	var contacts []Contact
	contacts, _ = kademlia.NodeLookup(network, target, false)
	return contacts
}

func (kademlia *Kademlia) findValue(network Network, hash string) string {
	var contacts []Contact
	var value []byte

	value = kademlia.Find(hash)
	if len(value) == 0 {
		contacts, value = kademlia.NodeLookup(network, hash, true)
	}

	if len(value) == 0 {
		// No value found
		return "No value with the hash " + hash + " was found. But " + strconv.Itoa(len(contacts)) + " close contacts was found."
	} else {
		// Return value
		return string(value)
	}
}

func (kademlia *Kademlia) Store(value []byte) string {
	key := HashValue(value)
	kademlia.HashTable[key] = value
	return key
}

func (kademlia *Kademlia) Find(key string) []byte {
	value := kademlia.HashTable[key]
	return value
}

func (kademlia *Kademlia) PingIp(network Network, destination string, sender Contact) bool {

	response := PingRPC(network, destination, sender)
	if response == "true" {
		fmt.Printf("Pong")
		return true
	} else {
		fmt.Printf("Could not ping node \n")
		return false
	}
}

func (kademlia *Kademlia) Ping(network Network, destination Contact, sender Contact) {

	response := PingRPC(network, destination.Address, sender)
	if response == "true" {
		fmt.Printf("Pong")
	} else {
		network.Node.RoutingTable.RemoveContact(destination)
		fmt.Printf("Could not ping node \n")
		return
	}
}
