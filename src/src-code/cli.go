package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Clilisten(kadnet Network, kademlia Kademlia, port int) {
	cmd := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd = scanner.Text()
		words := strings.Fields(cmd)
		answer := parse(words, kadnet, kademlia, port)
		fmt.Print("> " + answer + "\n")
		fmt.Printf(">")
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func parse(input []string, kadnet Network, kademlia Kademlia, port int) string {
	answer := ""
	switch input[0] {
	case "h":
		answer = "Ping should be like this: ping [ip address]\nPut should be like this: put [value]\nGet should be like this: get [hash]\nType exit to exit the node."
	case "ping":
		if len(input)>1 {
			dest := input[1] + ":" + strconv.Itoa(port)
			(*Kademlia).Ping(&kademlia, kadnet, dest, kademlia.Me)
		} else {
			answer = "Ping should be like this: ping [ip address]"
		}
	case "put":
		if len(input)>1 {
			value := []byte(input[1])
			(*Kademlia).PutCommand(&kademlia, kadnet, value)
		} else {
			answer = "Put should be like this: put [value]"
		}
	case "get":
		if len(input)>1 {
			hash := input[1]
			fmt.Printf("Send hash with lenght " + strconv.Itoa(len(hash)) + "\n")
			if len(hash) < 40 {
				answer = "Incorrect hash, the hash must be at least 40 chars"
			} else {
				(*Kademlia).GetCommand(&kademlia, kadnet, hash)
			}
		} else {
			answer = "Get should be like this: get [hash]"
		}
	case "exit":
		(*Kademlia).ExitCommand(&kademlia)

	//help commands for debugging
	case "routingtable":
		var contacts = kademlia.RoutingTable.FindClosestContacts(kademlia.Me.ID, 20)
		for _, contact := range contacts {
			answer += "Address: " + contact.Address + "\n>"
		}
	case "hashtable":
		answer = printHashTable(kademlia.HashTable)
	case "store":
		if len(input)>1 {
			value := []byte(input[1])
			answer = (*Kademlia).Store(&kademlia, value)
		} else {
			answer = "Store should be like this: store [value]"
		}
	case "ip":
		answer = kademlia.Me.Address
	case "storerpc":
		if len(input)>2 {
			value := []byte(input[1])
			destination := input[2] + ":" + strconv.Itoa(port)
			hash := StoreRPC(kadnet, destination, kademlia.Me, value)
			answer = hash
		} else {
			answer = "StoreRPC should be like this: storerpc [value] [ip]"
		}
	default:
		answer = "Unknown command " + input[0] + ", try again"
	}
	return answer
}

func printHashTable(hashtable map[string][]byte) string {
	answ := ""
	for k, v := range hashtable {
		answ += k + ": " + string(v) + "\n"
	}
	return answ
}

