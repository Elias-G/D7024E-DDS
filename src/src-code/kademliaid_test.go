package src

import (
	"reflect"
	"testing"
)

func TestKademliaID_CalcDistance(t *testing.T) {
	type args struct {
		target *KademliaID
	}
	tests := []struct {
		name       string
		kademliaID KademliaID
		args       args
		want       *KademliaID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kademliaID.CalcDistance(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKademliaID_Equals(t *testing.T) {
	type args struct {
		otherKademliaID *KademliaID
	}
	tests := []struct {
		name       string
		kademliaID KademliaID
		args       args
		want       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kademliaID.Equals(tt.args.otherKademliaID); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKademliaID_Less(t *testing.T) {
	type args struct {
		otherKademliaID *KademliaID
	}
	tests := []struct {
		name       string
		kademliaID KademliaID
		args       args
		want       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kademliaID.Less(tt.args.otherKademliaID); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKademliaID_String(t *testing.T) {
	tests := []struct {
		name       string
		kademliaID KademliaID
		want       string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kademliaID.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKademliaID(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *KademliaID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKademliaID(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKademliaID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRandomKademliaID(t *testing.T) {
	tests := []struct {
		name string
		want *KademliaID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomKademliaID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomKademliaID() = %v, want %v", got, tt.want)
			}
		})
	}
}