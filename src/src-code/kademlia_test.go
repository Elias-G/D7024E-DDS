package src

import (
	"fmt"
	"testing"
)

var me = CreateNode(5000, "10.0.0.3", NewRandomKademliaID())
var k = 5
var node = Kademlia{
	RoutingTable: *NewRoutingTable(me, k),
	Me:           me,
	K:            20,
	Alpha:        3,
	HashTable:    InitTable(),
	PingWait:     5,
}

func TestStore(t *testing.T) {
	send := []byte("piggy")
	want := HashValue(send)
	got := node.Store(send)

	if got != want {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", got, want)
	}
	if string(node.HashTable[got]) != string(send) {
		t.Errorf("TestStore does not store the value to its hashtable")
	}
}

func TestFind(t *testing.T) {
	want := []byte("piggy")
	hash := node.Store(want)

	got := node.Find(hash)

	if string(got) != string(want) {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", got, want)
	}
}

func TestPing_DeadNode(t *testing.T) {
	want := false

	var hashTable = InitTable()
	net := setup_net("10.0.0.3", NewKademliaID("0fda68927f2b2ff836f73578db0fa54c29f7fd92"), 5, 3, hashTable)
	id := NewKademliaID("0fda68927f2b2ff836f73578db0fa54c29f7fd92")
	contact := NewContact(id, "10.0.0.3")
	contact.CalcDistance(id)
	fmt.Print("Hi we are trying to ping")
	got := node.PingIp(net, "0.0.0.0:5000", contact)

	if got != want {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", got, want)
	}
}

func setup_net(ip string, id *KademliaID, k int, alpha int, hashTable map[string][]byte) Network {
	var me = CreateNode(5000, ip, id)
	me.CalcDistance(me.ID)
	var table = NewRoutingTable(me, k)
	var kademlia = &Kademlia{
		RoutingTable: *table,
		Me:           me,
		K:            k,
		Alpha:        alpha,
		HashTable:    hashTable,
		PingWait:     20000000000,
	}

	network := *NewNetwork(*kademlia)
	return network
}
