package main

import (
	"reflect"
	"src-code"
	"testing"
)

// Compares address of the node returned to the ip and port given
func TestCreateNode(t *testing.T) {
	type args struct {
		port int
		ip   string
	}
	tests := []struct {
		name string
		args args
		want string // Address
	}{
		// TODO: Add more test cases. Check more than address?
		{"CreateNode", args{1234, "1.2.3.4"}, "1.2.3.4:1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createNode(tt.args.port, tt.args.ip); !reflect.DeepEqual(got.Address, tt.want) {
				t.Errorf("Node address = %v, want %v", got.Address, tt.want)
			}
		})
	}
}

func Test_getIpAddress(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases. How to get IP to test???

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIpAddress(); got != tt.want {
				t.Errorf("getIpAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_networkJoin(t *testing.T) {
	type args struct {
		me       src.Contact
		rootNode src.Contact
		table    src.RoutingTable
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases. How to access rootnode?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
