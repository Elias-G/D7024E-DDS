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
		log.Printf("IP: " + ip)
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
		clilisten(ip, kadnet, *kademlia)
		//if arg == 2 then its a normal node to start
	} else if arg == "2" {
		var ip = getIpAddress()
		log.Printf("IP: " + ip)

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

		src.NetworkJoin(*kademlia, rootNode, *table, k)
		print(kademlia)
		//net.SendPingRequest(&rootNode, *kademlia)
		go kadnet.Listen(me.Address)
		clilisten(ip, kadnet, *kademlia)
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
		if i.Name == "eth0" {
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

func clilisten(ip string, kadnet src.Network, kademlia src.Kademlia) {
	cmd := ""
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd = scanner.Text()
		words := strings.Fields(cmd)
		parse(ip, words, kadnet, kademlia)
		//fmt.Println(reflect.TypeOf(words).String())
		fmt.Print("> ")
	}
	if scanner.Err() != nil {
		// handle error.
	}
}

func parse(ip string, input []string, kadnet src.Network, kademlia src.Kademlia) {
	switch input[0] {
	case "h":
		fmt.Print("This is help")
	case "ping":
		if(len(input)>2){
			sender := ip + ":" + input[2]
			dest := input[1] + ":" + input[2]
			//go kadnet.SendPingRequest(dest,sender)
			go (*src.Kademlia).Ping(&kademlia, kadnet, dest, sender)
		}
	default:
		fmt.Print("Try again")
	}
}
