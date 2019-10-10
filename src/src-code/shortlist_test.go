package src

import (
	"testing"
)

func TestShortList_Less(t *testing.T) {
	const nrOfContacts = 10
	shortList := generateShortList(nrOfContacts)

	dist1 := shortList.Contacts[0].ID.CalcDistance(shortList.Target)
	dist2 := shortList.Contacts[1].ID.CalcDistance(shortList.Target)

	want := dist1.Less(dist2)
	got := shortList.Less(0,1)

	if want != got {
		t.Errorf("Less = %v, want %v", got, want)
	}
}

func TestShortList_Sort(t *testing.T) {
	const nrOfContacts = 10
	shortList := generateShortList(nrOfContacts)

	shortList.Sort()
	for i := 0; i < len(shortList.Contacts)-1; i++ {
		if !shortList.Less(i, i+1) {
			t.Errorf("Sort not working")
		}
	}
}

func generateShortList(nrOfContacts int)(shortList ShortList) {
	target := NewRandomKademliaID()
	var contacts []Contact

	for i := 0; i < nrOfContacts; i++ {
		add := "Address" + string(i)
		contacts = append(contacts, NewContact(NewRandomKademliaID(), add))
	}

	return ShortList{target, contacts}
}