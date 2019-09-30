package src

import (
	"fmt"
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

func (kademlia *Kademlia) LookupData(kadnet src.Network, hash string) {
	// TODO returns data, if fails then it returns a list of adresses to send lookup to
	fmt.Print(hash)
	val, ok := kademlia.HashTable[hash]
	if ok {
		// Return the data
		fmt.Print(" success \n")
		fmt.Print(" this is the val: " + string(val) + " this was the hash: " + hash + "\n")
	} else {
		// Make RPC calls to the alpha nodes with the closest hashes in the routingtable
		//KadNetwork.
		//netsrc.SendPingRequest("10.0.0.3:5000", kademlia.Me.Address)
		fmt.Print(" fail \n")
		fmt.Print(" this is the val: " + string(val) + " this was the hash: " + hash + "\n")
	}
}

func (kademlia *Kademlia) Store(key string, value []byte) {
	// TODO
	kademlia.HashTable[key] = value
	fmt.Print(" this is the val: " + string(value) + "\n")
}
