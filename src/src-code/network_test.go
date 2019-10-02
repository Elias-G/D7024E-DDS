package src

import (
	"net"
	"reflect"
	"src-code/proto"
	"testing"
)

func TestNetworkJoin(t *testing.T) {
	type args struct {
		node     Kademlia
		rootNode Contact
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestNetwork_Listen(t *testing.T) {
	type fields struct {
		Node           Kademlia
		findNodeRespCh chan [] Contact
	}
	type args struct {
		address string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Node:           tt.fields.Node,
				findNodeRespCh: tt.fields.findNodeRespCh,
			}
		})
	}
}

func TestNetwork_NodeLookup(t *testing.T) {
	type fields struct {
		Node           Kademlia
		findNodeRespCh chan [] Contact
	}
	type args struct {
		id *KademliaID
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantContacts []Contact
		wantValue    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Node:           tt.fields.Node,
				findNodeRespCh: tt.fields.findNodeRespCh,
			}
			gotContacts, gotValue := network.NodeLookup(tt.args.id)
			if !reflect.DeepEqual(gotContacts, tt.wantContacts) {
				t.Errorf("NodeLookup() gotContacts = %v, want %v", gotContacts, tt.wantContacts)
			}
			if gotValue != tt.wantValue {
				t.Errorf("NodeLookup() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestNetwork_SendFindContactRequest(t *testing.T) {
	type fields struct {
		Node           Kademlia
		findNodeRespCh chan [] Contact
	}
	type args struct {
		contact     Contact
		kademliaObj Kademlia
		targetID    *KademliaID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Node:           tt.fields.Node,
				findNodeRespCh: tt.fields.findNodeRespCh,
			}
		})
	}
}

func TestNetwork_SendFindDataRequest(t *testing.T) {
	type fields struct {
		Node           Kademlia
		findNodeRespCh chan [] Contact
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Node:           tt.fields.Node,
				findNodeRespCh: tt.fields.findNodeRespCh,
			}
		})
	}
}

func TestNetwork_SendPingRequest(t *testing.T) {
	type fields struct {
		Node           Kademlia
		findNodeRespCh chan [] Contact
	}
	type args struct {
		destination string
		sender      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Node:           tt.fields.Node,
				findNodeRespCh: tt.fields.findNodeRespCh,
			}
		})
	}
}

func TestNetwork_SendStoreRequest(t *testing.T) {
	type fields struct {
		Node           Kademlia
		findNodeRespCh chan [] Contact
	}
	type args struct {
		contact     *Contact
		kademliaObj Kademlia
		data        []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Node:           tt.fields.Node,
				findNodeRespCh: tt.fields.findNodeRespCh,
			}
		})
	}
}

func TestNetwork_handleConnection(t *testing.T) {
	type fields struct {
		Node           Kademlia
		findNodeRespCh chan [] Contact
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network := &Network{
				Node:           tt.fields.Node,
				findNodeRespCh: tt.fields.findNodeRespCh,
			}
		})
	}
}

func TestNewNetwork(t *testing.T) {
	type args struct {
		node Kademlia
	}
	tests := []struct {
		name string
		args args
		want *Network
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNetwork(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatContactForReading(t *testing.T) {
	type args struct {
		contact kademlia.Contact
	}
	tests := []struct {
		name string
		args args
		want Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatContactForReading(tt.args.contact); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatContactForReading() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatContactForSending(t *testing.T) {
	type args struct {
		contact Contact
	}
	tests := []struct {
		name string
		args args
		want *kademlia.Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatContactForSending(tt.args.contact); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatContactForSending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatContactsForReading(t *testing.T) {
	type args struct {
		contacts []*kademlia.Contact
	}
	tests := []struct {
		name string
		args args
		want []Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatContactsForReading(tt.args.contacts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatContactsForReading() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatContactsForSending(t *testing.T) {
	type args struct {
		contacts []*Contact
	}
	tests := []struct {
		name string
		args args
		want []*kademlia.Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatContactsForSending(tt.args.contacts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatContactsForSending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatContactsForSending2(t *testing.T) {
	type args struct {
		contacts []Contact
	}
	tests := []struct {
		name string
		args args
		want []*kademlia.Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatContactsForSending2(tt.args.contacts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatContactsForSending2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFindNodeRequest(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want *kademlia.FindNodeRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFindNodeRequest(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFindNodeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFindNodeResponse(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want *kademlia.FindNodeResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFindNodeResponse(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFindNodeResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFindValueResponse(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want *kademlia.FindValueResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFindValueResponse(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFindValueResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readPingRequest(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want *kademlia.PingRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readPingRequest(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readPingRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readPingResponse(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want *kademlia.PingResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readPingResponse(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readPingResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readStoreRequest(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want *kademlia.StoreRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readStoreRequest(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readStoreRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readStoreResponse(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want *kademlia.StoreResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readStoreResponse(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readStoreResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sendData(t *testing.T) {
	type args struct {
		destination string
		dataToSend  []byte
		header      []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sendFindNodeResponse(t *testing.T) {
	type args struct {
		destination string
		sender      string
		ids         []Contact
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sendFindValueResponse(t *testing.T) {
	type args struct {
		destination string
		sender      string
		value       []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sendPingResponse(t *testing.T) {
	type args struct {
		destination string
		sender      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sendStoreResponse(t *testing.T) {
	type args struct {
		destination string
		sender      string
		value       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortContacts(t *testing.T) {
	type args struct {
		id       *KademliaID
		unsorted []Contact
		k        int
	}
	tests := []struct {
		name       string
		args       args
		wantSorted []Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSorted := sortContacts(tt.args.id, tt.args.unsorted, tt.args.k); !reflect.DeepEqual(gotSorted, tt.wantSorted) {
				t.Errorf("sortContacts() = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}