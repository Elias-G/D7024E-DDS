package src

import (
	"encoding/hex"
	kademliaProto "proto"
	"testing"
)

func Test_sendPingResponse(t *testing.T) {
	contactID := NewRandomKademliaID()
	rpcID := contactID.String()
	sender := GenerateContacts(1, contactID)[0]

	var i interface{} = sendPingResponse(rpcID, sender)

	_, ok := i.([]byte)
	if !ok {
		t.Errorf("sendPingResponse: Did not return []byte")
	}
}

func GenerateContacts(nrOfContacts int, id *KademliaID)(contacts []Contact) {
	for i := 0; i < nrOfContacts; i++ {
		contact := NewContact(NewRandomKademliaID(), "Address" + string(i))
		contact.CalcDistance(id)
		contacts = append(contacts, contact)
	}
	return contacts
}

func Test_sendFindNodeResponse(t *testing.T) {
	contactID := NewRandomKademliaID()
	rpcID := contactID.String()
	sender := GenerateContacts(1, contactID)[0]
	contacts := GenerateContacts(10, contactID)

	var i interface{} = sendFindNodeResponse(rpcID, sender, contacts)

	_, ok := i.([]byte)
	if !ok {
		t.Errorf("sendFindNodeResponse: Did not return []byte")
	}
}

func Test_sendFindValueResponse(t *testing.T) {
	contactID := NewRandomKademliaID()
	rpcID := contactID.String()
	sender := GenerateContacts(1, contactID)[0]
	contacts := GenerateContacts(10, contactID)
	value,_ := hex.DecodeString(rpcID)

	var i interface{} = sendFindValueResponse(rpcID, sender, value, contacts)

	_, ok := i.([]byte)
	if !ok {
		t.Errorf("sendFindValueResponse: Did not return []byte")
	}
}

func Test_sendStoreResponse(t *testing.T) {
	contactID := NewRandomKademliaID()
	rpcID := contactID.String()
	sender := GenerateContacts(1, contactID)[0]
	hash := "asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer"

	var i interface{} = sendStoreResponse(rpcID, sender, hash)

	_, ok := i.([]byte)
	if !ok {
		t.Errorf("sendStoreResponse: Did not return []byte")
	}
}

func Test_readPingRequest(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readPingRequest(message)

	_, ok := i.(*kademliaProto.PingRequest)
	if !ok {
		t.Errorf("sendPingRequest: Did not return *kademliaProto.PingRequest")
	}
}

func Test_readFindNodeRequest(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readFindNodeRequest(message)

	_, ok := i.(*kademliaProto.FindNodeRequest)
	if !ok {
		t.Errorf("readFindNodeRequest: Did not return *kademliaProto.FindNodeRequest")
	}
}

func Test_readFindValueRequest(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readFindValueRequest(message)

	_, ok := i.(*kademliaProto.FindValueRequest)
	if !ok {
		t.Errorf("readFindValueRequest: Did not return *kademliaProto.FindValueRequest")
	}
}

func Test_readStoreRequest(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readStoreRequest(message)

	_, ok := i.(*kademliaProto.StoreRequest)
	if !ok {
		t.Errorf("readStoreRequest: Did not return *kademliaProto.StoreRequest")
	}
}

func Test_readPingResponse(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readPingResponse(message)

	_, ok := i.(*kademliaProto.PingResponse)
	if !ok {
		t.Errorf("readStoreRequest: Did not return *kademliaProto.PingResponse")
	}
}

func Test_readFindNodeResponse(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readFindNodeResponse(message)

	_, ok := i.(*kademliaProto.FindNodeResponse)
	if !ok {
		t.Errorf("readFindNodeRequest: Did not return *kademliaProto.FindNodeResponse")
	}
}

func Test_readFindValueResponse(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readFindValueResponse(message)

	_, ok := i.(*kademliaProto.FindValueResponse)
	if !ok {
		t.Errorf("readFindValueResponse: Did not return *kademliaProto.FindValueResponse")
	}
}

func Test_readStoreResponse(t *testing.T) {
	message,_ := hex.DecodeString("asdfghjklzxcvbnmqwerasdfghjklzxcvbnmqwer")

	var i interface{} = readStoreResponse(message)

	_, ok := i.(*kademliaProto.StoreResponse)
	if !ok {
		t.Errorf("readStoreResponse: Did not return *kademliaProto.FindStoreResponse")
	}
}

func Test_formatContactsForRead(t *testing.T) {
	const nrOfContacts = 10
	id := NewRandomKademliaID()
	contacts := generateContactsForReading(nrOfContacts, id)

	var i interface{} = formatContactsForRead(contacts)

	_, ok := i.([]Contact)
	if !ok {
		t.Errorf("formatContactsForReading: Conversion failed")
	}
}

func generateContactsForReading(nrOfContacts int, id *KademliaID)(contacts []*kademliaProto.Contact) {
	for i := 0; i < nrOfContacts; i++ {
		contact := NewContact(NewRandomKademliaID(), "Address" + string(i))
		contact.CalcDistance(id)
		newContact := &kademliaProto.Contact{NodeId: contact.ID.String(), Address:contact.Address, Distance:contact.Distance.String()}
		contacts = append(contacts, newContact)
	}
	return contacts
}

func Test_formatContactForRead(t *testing.T) {
	contact := generateContactsForReading(1, NewRandomKademliaID())[0]
	var i interface {} = formatContactForRead(contact)

	_, ok := i.(Contact)
	if !ok {
		t.Errorf("formatContactForReading: Conversion failed")
	}
}