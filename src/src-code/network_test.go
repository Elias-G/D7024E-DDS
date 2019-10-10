package src

import (
	"strconv"
	"testing"
)

func Test_CreateNode(t *testing.T) {
	port := 1234
	ip := "1.2.3.4"
	id := NewRandomKademliaID()
	address := ip+":"+strconv.Itoa(port)

	contact := CreateNode(port, ip, id)
	if contact.Address != address {
		t.Errorf("CreateNode: wanted %v, got %v", address, contact.Address)
	} else if contact.ID != id {
		t.Errorf("CreateNode: wanted %v, got %v", id, contact.ID)
	}
}

func Test_dontAddRequester(t *testing.T) {
	contacts := GenerateContacts(10, NewRandomKademliaID())
	requester := contacts[8]
	// Requester is in list
	res := dontAddRequester(contacts, requester)
	for _,c := range res {
		if c.ID.Equals(requester.ID) {
			t.Errorf("dontAddRequester: requester is still in contact list")
		}
	}
	// Requester is in NOT list
	requester2 := GenerateContacts(1, NewRandomKademliaID())[0]
	res2 := dontAddRequester(contacts, requester2)
	for _,c := range res2 {
		if c.ID.Equals(requester.ID) {
			t.Errorf("dontAddRequester: requester is still in contact list")
		}
	}
}

func Test_NewNetwork(t *testing.T) {
	kad := Kademlia{
		RoutingTable: RoutingTable{},
		Me:           Contact{},
		K:            5,
		Alpha:        2,
		HashTable:    nil,
		PingWait:     0,
	}

	var i interface {} = NewNetwork(kad)

	_, ok := i.(*Network)
	if !ok {
		t.Errorf("NewNetwork: did not return *Network")
	}
}