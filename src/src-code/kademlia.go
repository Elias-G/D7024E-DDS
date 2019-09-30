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

func (kademlia *Kademlia) LookupData(hash string) string{
	// TODO returns data, if fails then it returns a list of adresses to send lookup to
	fmt.Print(hash)
	val, ok := kademlia.HashTable[hash]
	if ok {
		// Return the data
		fmt.Print(" successful lookup val: " + string(val) + " hash: " + hash )
		return string(val)
	} else {
		// Make RPC calls to the alpha nodes with the closest hashes in the routingtable
		//KadNetwork.
		//netsrc.SendPingRequest("10.0.0.3:5000", kademlia.Me.Address)
		fmt.Print(" fail hash: " + hash + "\n")
		return ""
	}
}

func (kademlia *Kademlia) LookupDataRoutingTable(hash string) []Contact {
	byteID := NewKademliaID(hash)
	contacts := kademlia.Table.FindClosestContacts(byteID , 20)
	return contacts
}

func (kademlia *Kademlia) Store(key string, value []byte) {
	// TODO
	kademlia.HashTable[key] = value
	fmt.Print(" Successful store val: " + string(value) + "\n")
}

func (kademlia *Kademlia) Ping(network Network, destination string, sender string) {
	// TODO
	var found = false
	timer := time.AfterFunc(time.Second * 5, func() {
		if found == false {
			//todo: remove from bucket
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
	//todo: move to front of bucket
	//find node contact in bucket and add it to front
}