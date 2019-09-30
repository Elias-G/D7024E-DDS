package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	src "src-code"
	"strconv"
	"strings"
)

var k = 20
var alpha = 3

func main() {
	arg := os.Args[1]
	//arg := "1"

	err := ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	// Initiate storage for node
	var hashTable = src.InitTable()

	//If arg==1 then its the rootnode that is suppose to start
	if arg == "1" {
		var ip = getIpAddress()
		var me = createNode(5000, ip)
		var table = src.NewRoutingTable(me)

		var kademlia = &src.Kademlia{
			Table:     *table,
			Me:        me,
			K:         k,
			Alpha:     1,
			HashTable: hashTable,
		}

		kadnet := *src.NewNetwork(*kademlia)

		print(kademlia)

		go kadnet.Listen(me.Address)
		clilisten(kademlia, kadnet)
		//if arg == 2 then its a normal node to start
	} else if arg == "2" {
		var ip = getIpAddress()

		var rootNode = createNode(5000, "10.0.0.3")

		var me = createNode(5000, ip)
		table := src.NewRoutingTable(me)

		var kademlia = &src.Kademlia{
			Table:     *table,
			Me:        me,
			K:         k,
			Alpha:     1,
			HashTable: hashTable,
		}

		kadnet := *src.NewNetwork(*kademlia)

		src.NetworkJoin(*kademlia, rootNode)
		print(kademlia)
		//net.SendPingRequest(&rootNode, *kademlia)
		go kadnet.Listen(me.Address)
		clilisten(kademlia, kadnet)
	} else {
		fmt.Print("Choose to be a leader(1) or a follower(2)")
	}

}

func getIpAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("interface error", err)
	}

	for _, i := range ifaces {
		if i.Name == "eth1" {
			addrs, err := i.Addrs()
			if err != nil {
				log.Fatal("interface error", err)
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				fmt.Printf(ip.String())
				return ip.String()
			}
		}
	}

	return ""
}

func createNode(port int, ip string) src.Contact {
	id := src.NewRandomKademliaID()
	address := ip + ":" + strconv.Itoa(port)
	me := src.NewContact(id, address)
	return me
}

func clilisten(node *src.Kademlia, kadnet src.Network) {
	cmd := ""
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd = scanner.Text()
		words := strings.Fields(cmd)
		parse(node, words, kadnet)
		//fmt.Println(reflect.TypeOf(words).String())
		fmt.Print("> ")
	}
	if scanner.Err() != nil {
		// handle error.
	}
}

func parse(node *src.Kademlia, input []string, kadnet src.Network) {
	switch input[0] {
	case "h":
		fmt.Print("This is helpful")
		fmt.Print(node.Me.Address)
	case "ping":
		dest := input[1]
		ipport := dest + ":5000"
		go kadnet.SendPingRequest(ipport, node.Me.Address)

	case "get":
		hash := input[1]
		data  := node.LookupData(hash)
		if data == "" {
			contacts := node.LookupDataRoutingTable(hash)
			fmt.Print(len(contacts))
			if len(contacts) == 0{
				fmt.Print("we found no contacts...?")	
			}else{
				go kadnet.SendFindValueRequest(&contacts[0], hash)
			}
			//go kadnet.SendFindValueRequest(contacts[1])
			//go kadnet.SendFindValueRequest(contacts[2])

		}else{
			fmt.Print(data)
		}

	case "store":
		val := input[1]
		hash := input[2]
		go node.Store(hash, []byte(val))

	case "ip":
		fmt.Print(node.Me.Address)

	case "RTsize":
		hash := input[1]
		contacts := node.LookupDataRoutingTable(hash)
		fmt.Print(len(contacts))
	default:
		fmt.Print("Try again")
	}
}
