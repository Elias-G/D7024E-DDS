package src

import (
	"fmt"
	//"src-code"
	//netsrc "src-code/network"
)

type Kademlia struct {
	Table     RoutingTable
	Me        Contact
	K         int
	Alpha     int
	HashTable map[string][]byte
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
