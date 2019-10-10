package src

import (
	"encoding/hex"
	"fmt"
	"strconv"
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

func TestNewKademliaID(t *testing.T) {
	bytes, _ := hex.DecodeString("asdfghjklzxcvbnmasdfasdfghjklzxcvbnmasdf")
	key := HashValue(bytes)
	NewKademliaID(key)
}


func TestNewKademliaID2(t *testing.T) {
bytes, _ := hex.DecodeString("piggy")
key := HashValue(bytes)
fmt.Printf(strconv.Itoa(len(key)))
NewKademliaID(key)
}