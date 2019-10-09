package src

import (
	"testing"
)

func TestSortContacts(t *testing.T) {
	targetID := NewRandomKademliaID()
	contacts := GenerateContacts(10, targetID)

	got := sortContacts(targetID, contacts)

	for i, contact := range got {
		if i > 0 && contact.Distance.Less(got[i-1].Distance) {
			t.Errorf("Sorting failed")
		}
	}
}

func TestUpdateShortList(t *testing.T) {
	lenContacts := 10
	lenShortList := 10
	nrOfProbed := 4

	targetID := NewRandomKademliaID()
	contacts := GenerateContacts(lenContacts, targetID)
	shortList := GenerateContacts(lenShortList, targetID)

	var probed []Contact
	for i := 0; i < nrOfProbed; i++ {
		probed = append(probed, contacts[i*2])
	}

	got := updateShortList(contacts, targetID, shortList, probed)

	for _,contact := range got {
		for _, used := range probed {
			if contact.ID.Equals(used.ID) {
				t.Errorf("updateShortList failded")
			}
		}
	}
	wantedLen := lenContacts + lenShortList - nrOfProbed
	if len(got) != wantedLen {
		t.Errorf("updateShortList failded, wrong length, wanted %v, got %v", wantedLen, len(got))
	}
}
