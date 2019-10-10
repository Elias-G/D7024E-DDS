package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Clilisten(ip string, kadnet Network, kademlia Kademlia, port int) {
	cmd := ""
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd = scanner.Text()
		words := strings.Fields(cmd)
		parse(words, kadnet, kademlia, port)
		fmt.Print("> ")
	}
	if scanner.Err() != nil {
		// handle error. //todo
	}
}

func parse(input []string, kadnet Network, kademlia Kademlia, port int) {
	switch input[0] {
	case "h":
		fmt.Print("Ping should be like this: ping [ip address]\nPut should be like this: put [value]\nGet should be like this: get [hash]\nType exit to exit the node.\n")
	case "ping":
		if len(input)>1 {
			dest := input[1] + ":" + strconv.Itoa(port)
			go (*Kademlia).Ping(&kademlia, kadnet, dest, kademlia.Me)
		}else {
			fmt.Printf("Ping should be like this: ping [ip address]\n")
		}
	case "put":
		if len(input)>1 {
			value := []byte(input[1])
			go (*Kademlia).PutCommand(&kademlia, kadnet, value)
		}else {
			fmt.Printf("Put should be like this: put [value]\n")
		}
	case "get":
		if len(input)>1 {
			hash := input[1]
			fmt.Printf("Send hash with lenght " + strconv.Itoa(len(hash)) + "\n")
			if len(hash) < 40 {
				fmt.Printf("Incorrect hash, the hash must be at least 40 chars\n")
			} else {
				go (*Kademlia).GetCommand(&kademlia, kadnet, hash)
			}
		}else {
			fmt.Printf("Get should be like this: get [hash]\n")
		}
	case "exit":
		go (*Kademlia).ExitCommand(&kademlia, kadnet)

	//help commands for debugging
	case "routingtable":
		var contacts = kademlia.RoutingTable.FindClosestContacts(kademlia.Me.ID, 20)
		fmt.Print(len(contacts))
		for _, contact := range contacts {
			fmt.Printf("Address: " + contact.Address + "\n")
		}
	case "hashtable":
		fmt.Print(kademlia.HashTable)
		fmt.Printf("\n")
	case "store":
		(*Kademlia).Store(&kademlia,[]byte("piggy"))
	case "ip":
		fmt.Printf(kademlia.Me.Address + "\n")
	default:
		fmt.Print("Unknown command " + input[0] + ", try again\n")
	}
}
