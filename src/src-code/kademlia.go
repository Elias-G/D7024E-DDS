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

func (kademlia *Kademlia) Ping(network Network, destination string, sender string) {
	// TODO
	network.SendPingRequest(destination, sender)
	//wait for response
	//if response: move to front of bucket
	//else: remove from bucket
}