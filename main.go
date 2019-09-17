package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"./src"
)

var k = 20

func main() {
	var ip = getIpAddress()

	var rootNode = createNode(5000, "10.0.0.3")

	var me = createNode(5000, ip)
	table := src.NewRoutingTable(me)

	var kademlia = &src.Kademlia{
		Table: *table,
		Me:    me,
		K:     k,
		Alpha: 1,
	}
	src.NetworkJoin(me, rootNode, *table, k)

	print(kademlia)

	var kadnet = src.Network{}

	//net.SendPingRequest(&rootNode, *kademlia)
	clilisten()
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

func clilisten() {
	cmd := ""
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd = scanner.Text()
		words := strings.Fields(cmd)
		parse(words)
		//fmt.Println(reflect.TypeOf(words).String())
		fmt.Print("> ")
	}
	if scanner.Err() != nil {
		// handle error.
	}
}

func parse(input []string) {
	switch input[0] {
	case "h":
		fmt.Print("This is help")
	case "ping":
		ip := input[1]
		go kadnet.SendPingMessage(&rootNode, ip)
	default:
		fmt.Print("Try again")
	}
}
