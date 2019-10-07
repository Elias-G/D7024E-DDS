package src

import (
	"fmt"
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

/*
CLI commands
 */
func (kademlia *Kademlia) PutCommand(network Network, value []byte) {
	// TODO
	fmt.Printf("This should return a hash!\n")
}

func (kademlia *Kademlia) GetCommand(network Network, hash string) {
	// TODO
	fmt.Printf("This should return a value!\n")
}

func (kademlia *Kademlia) ExitCommand(network Network) {
	// TODO
	fmt.Printf("This should Exit!\n")
}


/*
RPCs
 */
func (kademlia *Kademlia) findNode(network Network, target *Contact) {
	var contacts []Contact
	contacts,_ = kademlia.NodeLookup(network, target, "", false)

	if contacts[0].ID.Equals(target.ID) {
		// Found what you were looking for! :D
	} else {
		// Close enough
	}
}

func (kademlia *Kademlia) findValue(network Network, hash string) string {
	var contacts []Contact
	var value string
	contacts, value = kademlia.NodeLookup(network, nil, hash, true)

	if value == "" {
		// No value found
		return "No value with the hash " + hash + " was found. But " + string(len(contacts)) + " close contacts was found."
	} else {
		// Return value
		return value
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
			network.Node.RoutingTable.RemoveContact(sender)//todo: implement remove from bucket and make sender a contact
			//find node contact in bucket and remove it
			fmt.Printf("Could not ping node \n")
			return
		}
	})

	response := PingRPC(network, destination, sender)

	found=true //if this code is reached a response came back and node is alive
	timer.Stop() //then timer can be stopped
	fmt.Printf(response) //print the result todo: should this be printed?
	network.Node.RoutingTable.AddContact(sender) //todo: make sender to contact, to add to front of bucket
}