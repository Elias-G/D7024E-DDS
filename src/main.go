package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	src "src-code"
)

//program variables
var k = 3
var alpha = 1
var rootId = src.NewKademliaID("0fda68927f2b2ff836f73578db0fa54c29f7fd92")
var port = 5000

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
		var ip = src.GetIpAddress()
		var me = src.CreateNode(5000, ip, rootId)
		me.CalcDistance(me.ID)
		log.Printf("IP: " + ip + " kademlia id: " + me.ID.String())
		var table = src.NewRoutingTable(me)

		var kademlia = &src.Kademlia{
			RoutingTable:     *table,
			Me:        me,
			K:         k,
			Alpha:     alpha,
			HashTable: hashTable,
			PingWait:  20000000000,
		}

		network := *src.NewNetwork(*kademlia)

		print(kademlia)

		go network.Listen(me.Address)
		src.Clilisten(network, *kademlia, port)
		//if arg == 2 then its a normal node to start
	} else if arg == "2" {
		var ip = src.GetIpAddress()

		var rootNode = src.CreateNode(5000, "10.0.0.3", rootId)
		var me = src.CreateNode(5000, ip, src.NewRandomKademliaID())
		me.CalcDistance(me.ID)
		rootNode.CalcDistance(me.ID)


		log.Printf("IP: " + ip + " kademlia id: " + me.ID.String())
		table := src.NewRoutingTable(me)

		var kademlia = &src.Kademlia{
			RoutingTable:     *table,
			Me:        me,
			K:         k,
			Alpha:     alpha,
			HashTable: hashTable,
		}

		network := *src.NewNetwork(*kademlia)

		go network.Listen(me.Address)

		network.NetworkJoin(*kademlia, rootNode)


		src.Clilisten(network, *kademlia, port)
	} else {
		fmt.Print("Choose to be a leader(1) or a follower(2)")
	}

}
