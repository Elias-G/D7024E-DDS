package main

import (
	"./src"
	"fmt"
	"log"
	"net"
	"strconv"
)

var k2 = 20

func main() {
	var ip = getIpAddress2()
	var me = createNode2(5000, ip)
	var table = src.NewRoutingTable(me)

	var kademlia = &src.Kademlia{
		Table: *table,
		Me:    me,
		K:     k2,
		Alpha: 1,
	}

	print(kademlia)

	src.Listen(me.Address)
}

func createNode2(port int, ip string) src.Contact {
	id := src.NewRandomKademliaID()
	address := ip + ":" + strconv.Itoa(port)
	me := src.NewContact(id, address)
	return me
}

func getIpAddress2() string {
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
