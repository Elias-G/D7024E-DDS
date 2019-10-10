package src

import (
	"testing"
)

var me = CreateNode(5000, "10.0.0.3", NewRandomKademliaID())
var node = Kademlia{
	RoutingTable: *NewRoutingTable(me),
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