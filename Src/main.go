package main

import (
	"Src"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var k = 20
var kadnet = Src.Network{}

func main() {
	//arg := os.Args[1]
	arg := "1"

	err := ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}


	//If arg==1 then its the rootnode that is suppose to start
	if arg == "1" {
		var ip = getIpAddress()
		var me = createNode(5000, ip)
		var table = Src.NewRoutingTable(me)

		var kademlia = &Src.Kademlia{
			Table: *table,
			Me:    me,
			K:     k,
			Alpha: 1,
		}

		print(kademlia)

		Src.Listen(me.Address)
		//if arg == 2 then its a normal node to start
	} else if arg == "2" {
		var ip = getIpAddress()

		var rootNode = createNode(5000, "10.0.0.3")

		var me = createNode(5000, ip)
		table := Src.NewRoutingTable(me)

		var kademlia = &Src.Kademlia{
			Table: *table,
			Me:    me,
			K:     k,
			Alpha: 1,
		}
		Src.NetworkJoin(me, rootNode, *table, k)
		print(kademlia)
		//net.SendPingRequest(&rootNode, *kademlia)
		clilisten(ip)
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

func createNode(port int, ip string) Src.Contact {
	id := Src.NewRandomKademliaID()
	address := ip + ":" + strconv.Itoa(port)
	me := Src.NewContact(id, address)
	return me
}

func clilisten(ip string) {
	cmd := ""
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd = scanner.Text()
		words := strings.Fields(cmd)
		parse(ip, words)
		//fmt.Println(reflect.TypeOf(words).String())
		fmt.Print("> ")
	}
	if scanner.Err() != nil {
		// handle error.
	}
}

func parse(ip string, input []string) {
	switch input[0] {
	case "h":
		fmt.Print("This is help")
	case "ping":
		dest := input[1]
		ipport := dest + ":5000"
		go kadnet.SendPingRequest(ip, ipport)
	default:
		fmt.Print("Try again")
	}
}
