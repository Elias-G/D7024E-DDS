package src

import (
	"testing"
)

var network = NewNetwork(node)

func TestParse_h(t *testing.T) {
	input := []string{"h"}
	got := parse(input, *network, node, 5000)

	want := "Ping should be like this: pingip [ip address]\nPut should be like this: put [value]\nGet should be like this: get [hash]\nType exit to exit the node."
	if got != want {
		t.Errorf("TestParse_h = %v, want %v", got, want)
	}
}

func TestParse_ping(t *testing.T) {
	input := []string{"pingip"}
	got := parse(input, *network, node, 5000)

	want := "Pingip should be like this: pingip [ip address]"
	if got != want {
		t.Errorf("TestParse_pingip = %v, want %v", got, want)
	}
}

func TestParse_put(t *testing.T) {
	input := []string{"put"}
	got := parse(input, *network, node, 5000)

	want := "Put should be like this: put [value]"
	if got != want {
		t.Errorf("TestParse_put = %v, want %v", got, want)
	}
}

func TestParse_get(t *testing.T) {
	input := []string{"get"}
	got := parse(input, *network, node, 5000)
	want := "Get should be like this: get [hash]"

	if got != want {
		t.Errorf("TestParse_get = %v, want %v", got, want)
	}

	input = []string{"get", "sdfghj"}
	got = parse(input, *network, node, 5000)
	want = "Incorrect hash, the hash must be at least 40 chars"

	if got != want {
		t.Errorf("TestParse_get = %v, want %v", got, want)
	}
}

func TestParse_ip(t *testing.T) {
	input := []string{"ip"}
	got := parse(input, *network, node, 5000)
	want := node.Me.Address

	if got != want {
		t.Errorf("TestParse_ip = %v, want %v", got, want)
	}
}

func TestParse_default(t *testing.T) {
	input := []string{"unknown"}
	got := parse(input, *network, node, 5000)
	want := "Unknown command " + input[0] + ", try again"

	if got != want {
		t.Errorf("TestParse_default = %v, want %v", got, want)
	}
}

func TestParse_routingtable(t *testing.T) {
	contacts := GenerateContacts(1, NewRandomKademliaID())
	node.RoutingTable.AddContact(contacts[0])

	input := []string{"routingtable"}
	got := parse(input, *network, node, 5000)
	want := "ID: " + contacts[0].ID.String() + " Address: " + contacts[0].Address + "\n>"
	if got != want {
		t.Errorf("TestParse_routingtable = %v, want %v", got, want)
	}
}

func TestParse_store(t *testing.T) {
	node.HashTable = InitTable()
	value := []byte("piggy")
	hash := HashValue(value)

	input := []string{"store", string(value)}
	got := parse(input, *network, node, 5000)
	want := hash

	if got != want {
		t.Errorf("TestParse_store = %v, want %v", got, want)
	}

	input = []string{"store"}
	got = parse(input, *network, node, 5000)
	want = "Store should be like this: store [value]"

	if got != want {
		t.Errorf("TestParse_store = %v, want %v", got, want)
	}
}

func TestPrintHashTable(t *testing.T) {
	node.HashTable = InitTable()
	node.HashTable["test1"] = []byte("123")

	got := printHashTable(node.HashTable)
	want := "test1: 123\n"

	if got != want {
		t.Errorf("TestPrintHashTable = %v, want %v", got, want)
	}
}

func TestParse_hashtable(t *testing.T) {
	node.HashTable = InitTable()
	key := "test"
	value := []byte("123")
	node.HashTable[key] = value

	input := []string{"hashtable"}
	got := parse(input, *network, node, 5000)

	want := key + ": " + string(value) + "\n"
	if got != want {
		t.Errorf("TestParse_hashtable = %v, want %v", got, want)
	}
}
