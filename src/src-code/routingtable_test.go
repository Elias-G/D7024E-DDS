package src

import (
	"fmt"
	"testing"
)

func TestRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"), k)

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

func TestRoutingTable_Remove(t *testing.T) {
	node := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	contact1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000001"), "localhost:8001")
	contact2 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000002"), "localhost:8002")
	contact3 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000003"), "localhost:8003")
	rt := NewRoutingTable(node, k)
	rt.AddContact(contact1)
	rt.AddContact(contact2)
	rt.AddContact(contact3)
	rt.RemoveContact(contact1)
	fmt.Print(rt.String())
	rtwant := NewRoutingTable(node, k)
	rtwant.AddContact(contact2)
	rtwant.AddContact(contact3)
	if rt.String() != rtwant.String() {
		t.Errorf("got %v, want %v", rt, rtwant)
	}
}

func TestRoutingTable_String(t *testing.T) {
	node := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000")
	contact1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000001"), "localhost:8001")
	contact2 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000002"), "localhost:8002")
	contact3 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000003"), "localhost:8003")
	rt := NewRoutingTable(node, k)
	rt.AddContact(contact1)
	rt.AddContact(contact2)
	rt.AddContact(contact3)
	rtwant := "\nAddress: localhost:8001\nAddress: localhost:8002\nAddress: localhost:8003\n"
	if rt.String() != rtwant {
		t.Errorf("got %v, want %v", rt, rtwant)
	}
}
