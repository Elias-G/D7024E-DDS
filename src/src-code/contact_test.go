package src

import (
	"testing"
)

func TestNewContact(t *testing.T) {
	type args struct {
		ID       *KademliaID
		Address  string
		Distance *KademliaID
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"NewContact", args{NewRandomKademliaID(), "1.0.0.1:1001", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewContact(tt.args.ID, tt.args.Address)
			if got.Address != tt.args.Address {
				t.Errorf("Address = %v, want %v", got.Address, tt.args.Address)
			}
			if got.ID != tt.args.ID {
				t.Errorf("ID = %v, want %v", got.ID, tt.args.ID)
			}
			// TODO: Test distance
			/*if assert.Nil(t, got.Distance) {
				t.Errorf("Distance = %v, want %v", got.Distance, tt.args.ID)
			}*/
		})
	}
}
