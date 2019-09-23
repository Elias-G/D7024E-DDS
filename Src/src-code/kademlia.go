package src

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
	// TODO
}

func (kademlia *Kademlia) Store(key string, value []byte) {
	// TODO
	kademlia.HashTable[key] = value
}
