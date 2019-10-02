package src

import (
	"fmt"
	"time"
)

type Kademlia struct {
	Table     RoutingTable
	Me        Contact
	K         int
	Alpha     int
	HashTable map[string][]byte
	PingWait  time.Duration
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(key string, value []byte) {
	// TODO
	kademlia.HashTable[key] = value
}

func (kademlia *Kademlia) Ping(network Network, destination string, sender string) {
	var found = false
	timer := time.AfterFunc(time.Second * 5, func() {
		if found == false {
			//network.Node.Table.RemoveContact(sender)//todo: implement remove from bucket and make sender a contact
			//find node contact in bucket and remove it
			fmt.Printf("Could not ping node \n")
			return
		}
	})
	network.SendPingRequest(destination, sender)

	response := <-network.pingRespCh
	found=true
	timer.Stop()
	fmt.Printf("Response: " + response.GetResponse() + " from sender: " + response.GetSender() + "\n")
	//network.Node.Table.AddContact(sender) //todo: make sender to contact, to add to front of bucket
	//find node contact in bucket and add it to front
}