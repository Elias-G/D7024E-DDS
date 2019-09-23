package src

import (
	"fmt")

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

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO returns data, if fails then it returns a list of adresses to send lookup to
	fmt.Print(hash)
	val, ok := kademlia.HashTable[hash]
	if(ok){
		// Return the data
		fmt.Print(" success \n")
		fmt.Print(" this is the val: " + string(val) +" this was the hash: " + hash + "\n")
	}else{
		// Make RPC calls to the alpha nodes with the closest hashes in the routingtable
		fmt.Print(" fail \n")
		fmt.Print(" this is the val: " + string(val) +" this was the hash: " + hash +"\n")
	}
}

func (kademlia *Kademlia) Store(key string, value []byte) {
	// TODO
	kademlia.HashTable[key] = value
	fmt.Print(" this is the val: " + string(value) +"\n")
}
