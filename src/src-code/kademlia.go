package src

import (
	"fmt"
	"os"
	"strconv"
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

	fmt.Printf("Nodes found: " + strconv.Itoa(len(nodes)) + "\n")

	for _, node := range nodes {
		StoreRPC(network, node.Address, kademlia.Me, value)
		fmt.Printf("Sending " + string(value) + " to " + node.Address + " from " + kademlia.Me.Address + "\n")
	}

	fmt.Printf(hash + "\n")
	fmt.Printf("Recieved hash with lenght " + strconv.Itoa(len(hash)) + "\n")
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
	contacts,_ = kademlia.NodeLookup(network, target, false)
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