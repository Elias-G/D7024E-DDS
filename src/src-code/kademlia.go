package src

import (
	"fmt"
	"os"
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
	hash := HashValue(value)
	nodes := kademlia.findNode(network, hash)

	for _, node := range nodes {
		StoreRPC(network, node.Address, kademlia.Me, value)
	}

	fmt.Printf(hash + "\n")
}

func (kademlia *Kademlia) GetCommand(network Network, hash string) {
	value := kademlia.findValue(network, hash)
	fmt.Printf(value + "\n")
}

func (kademlia *Kademlia) ExitCommand(network Network) {
	os.Exit(1)
}


/*
RPCs
 */
func (kademlia *Kademlia) findNode(network Network, target string) []Contact {
	var contacts []Contact
	contacts,_ = kademlia.NodeLookup(network, nil, target, false)
	return contacts
}

func (kademlia *Kademlia) findValue(network Network, hash string) string {
	var contacts []Contact
	var value string
	contacts, value = kademlia.NodeLookup(network, nil, hash, true)

	if value == "" {
		return "No value with the hash " + hash + " was found. But " + string(len(contacts)) + " close contacts was found."
	} else {
		return value
	}
}

func (kademlia *Kademlia) Store(value []byte) string {
	key := HashValue(value)
	kademlia.HashTable[key] = value
	return key
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

	response := PingRPC(network, destination, sender)

	found=true //if this code is reached a response came back and node is alive
	timer.Stop() //then timer can be stopped
	fmt.Printf(response) //print the result todo: should this be printed?
	//network.Node.Table.AddContact(sender) //todo: make sender to contact, to add to front of bucket
}