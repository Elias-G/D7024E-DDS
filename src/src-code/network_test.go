package src

import (
	"encoding/hex"
	kademlia "proto"
	//"reflect"
	"testing"
)

func GenerateContacts(nrOfContacts int, id *KademliaID)(contacts []Contact) {
	for i := 0; i < nrOfContacts; i++ {
		contact := NewContact(NewRandomKademliaID(), "Address" + string(i))
		contact.CalcDistance(id)
		contacts = append(contacts, contact)
	}
	return contacts
}

func generateContactsForSending(nrOfContacts int, id *KademliaID)(contacts []*Contact) {
	for i := 0; i < nrOfContacts; i++ {
		contact := NewContact(NewRandomKademliaID(), "Address" + string(i))
		contact.CalcDistance(id)
		contacts = append(contacts, &contact)
	}
	return contacts
}

func generateContactsForSending2(nrOfContacts int, id *KademliaID)(contacts []Contact) {
	for i := 0; i < nrOfContacts; i++ {
		contact := NewContact(NewRandomKademliaID(), "Address" + string(i))
		contact.CalcDistance(id)
		contacts = append(contacts, contact)
	}
	return contacts
}

func generateContactsForReading(nrOfContacts int, id *KademliaID)(contacts []*kademlia.Contact) {
	for i := 0; i < nrOfContacts; i++ {
		contact := NewContact(NewRandomKademliaID(), "Address" + string(i))
		contact.CalcDistance(id)
		newContact := &kademlia.Contact{NodeId: contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()}
		contacts = append(contacts, newContact)
	}
	return contacts
}

func generateMessage()(message []byte) {
	message,_ = hex.DecodeString("MessageToByte")
	return message
}

func TestNetworkJoin(t *testing.T) {

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
			//if got := NewNetwork(tt.args.node, make(chan *kademlia.PingResponse), make(chan []Contact)); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewNetwork() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func Test_formatContactForSending(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "Address")
	id := NewRandomKademliaID()
	contact.CalcDistance(id)

	var i interface{} = formatContactForSend(contact)

	_, ok := i.(*kademlia.Contact)
	if !ok {
		t.Errorf("formatContactForSending: Conversion failed")
	}
}

func Test_formatContactsForReading(t *testing.T) {
	const nrOfContacts = 10
	id := NewRandomKademliaID()
	contacts := generateContactsForReading(nrOfContacts, id)

	var i interface{} = formatContactsForRead(contacts)

	_, ok := i.([]Contact)
	if !ok {
		t.Errorf("formatContactsForReading: Conversion failed")
	}
}

func Test_formatContactsForSending(t *testing.T) {
	const nrOfContacts = 10
	id := NewRandomKademliaID()
	contacts := generateContactsForSending(nrOfContacts, id)
	var i interface{} = formatContactsForSend(contacts)

	_, ok := i.([]*kademlia.Contact)
	if !ok {
		t.Errorf("formatContactsForSending: Conversion failed")
	}
}

func Test_formatContactsForSending2(t *testing.T) {
	const nrOfContacts = 10
	id := NewRandomKademliaID()
	contacts := generateContactsForSending2(nrOfContacts, id)
	var i interface{} = formatContactsForSend2(contacts)

	_, ok := i.([]*kademlia.Contact)
	if !ok {
		t.Errorf("formatContactsForSending2: Conversion failed")
	}
}

func Test_readFindNodeRequest(t *testing.T) {
	message := generateMessage()

	var i interface{} = readFindNodeRequest(message)

	_, ok := i.(*kademlia.FindNodeRequest)
	if !ok {
		t.Errorf("readFindNodeRequest: Read failed")
	}
}

func Test_readFindNodeResponse(t *testing.T) {
	message := generateMessage()

	var i interface{} = readFindNodeResponse(message)

	_, ok := i.(*kademlia.FindNodeResponse)
	if !ok {
		t.Errorf("readFindNodeResponse: Read failed")
	}
}

func Test_readFindValueResponse(t *testing.T) {
	message := generateMessage()

	var i interface{} = readFindValueResponse(message)

	_, ok := i.(*kademlia.FindValueResponse)
	if !ok {
		t.Errorf("readFindValueResponse: Read failed")
	}
}

func Test_readPingRequest(t *testing.T) {
	message := generateMessage()

	var i interface{} = readPingRequest(message)

	_, ok := i.(*kademlia.PingRequest)
	if !ok {
		t.Errorf("readPingRequest: Read failed")
	}
}

func Test_readPingResponse(t *testing.T) {
	message := generateMessage()

	var i interface{} = readPingResponse(message)

	_, ok := i.(*kademlia.PingResponse)
	if !ok {
		t.Errorf("readPingResponse: Read failed")
	}
}

func Test_readStoreRequest(t *testing.T) {
	message := generateMessage()

	var i interface{} = readStoreRequest(message)

	_, ok := i.(*kademlia.StoreRequest)
	if !ok {
		t.Errorf("readStoreRequest: Read failed")
	}
}

func Test_readStoreResponse(t *testing.T) {
	message := generateMessage()

	var i interface{} = readStoreResponse(message)

	_, ok := i.(*kademlia.StoreResponse)
	if !ok {
		t.Errorf("readStoreResponse: Read failed")
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
	const nrOfContacts = 10
	const k = 8
	id := NewRandomKademliaID()
	contacts := GenerateContacts(nrOfContacts, id)

	got := sortContacts(id, contacts)

	var i interface{} = got

	_, ok := i.([]Contact)
	if !ok {
		t.Errorf("sortContacts: Did not return type []Contact")
	}

	if len(got) != k {
		t.Errorf("sortContacts: Did not return k contacts")
	}
}