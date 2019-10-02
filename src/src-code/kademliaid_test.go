package src

import (
	"testing"
)

func TestKademliaID_Equals(t *testing.T) {
	id := NewRandomKademliaID()
	got := id.Equals(id)
	want := true

	if got != want {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", got, want)
	}
}

