package main

import (
	"./src"
	"fmt"
	"log"
	"net"
	"strconv"
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

	var net = src.Network{}
	net.SendPingRequest(rootNode.Address, kademlia.Me.Address)
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
