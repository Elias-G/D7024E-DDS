package src

import "fmt"

// Sends out alpha RPCs for FindNode/FindValue and returns k closest contacts or value if found
// TODO: Parallel requests, Timing?
func (kademlia *Kademlia) NodeLookup(network Network, target string, findValue bool)(contacts []Contact, value []byte) {
	var table = kademlia.RoutingTable
	var alpha = kademlia.Alpha
	//var k = kademlia.K
	var shortList []Contact
	var probed []Contact // keep track of contacts that have already been probed
	var noVal []byte // If no value was found or searched for, return empty byte array
	var targetID *KademliaID

	targetID = NewKademliaID(target)
	shortList = table.FindClosestContacts(targetID, alpha)
	fmt.Printf("SHORTLIST, findClosestContact, nodeLookup: " + printContacts(shortList) + "\n")
	//var closestNode = shortList[0].ID // current closest node to target //todo: uncomment when iterative is used

	// TODO: Parallel?
	for i := 0; i < alpha; i++ {
		var contact = shortList[i]
		if findValue {
			shortList, probed, value = kademlia.sendFindValueRPCs(network, &contact, target, shortList, probed)
			if len(value) != 0 {
				return nil, value
			}
		} else {
			shortList, probed = kademlia.sendFindNodeRPCs(network, contact, targetID, shortList, probed)
			fmt.Printf("SHORTLIST NEW, sendFindNodeRPCs, nodeLookup: " + printContacts(shortList) + "\n")
		}
	}

	//contacts, value =kademlia.iterativeLookup(network, shortList, probed, target, *targetID, closestNode, findValue, value)

	contacts=shortList

	fmt.Printf("CONTACTS, nodeLookup: " + printContacts(contacts) + "\n")

	return contacts, noVal
}

/*func (kademlia *Kademlia) iterativeLookup(network Network, shortList []Contact, probed []Contact, target string, targetID KademliaID, closestNode *KademliaID, findValue bool, value []byte) ([]Contact, []byte) {
	// While recent responses are closer than closestNode, send new RPCs
	for shortList[0].ID.CalcDistance(&targetID).Less(closestNode.CalcDistance(&targetID)) {
		closestNode = shortList[0].ID
		// if less than k nodes has been successfully probed
		if len(probed) < kademlia.K { //todo: move if to while
			for i := 0; i < kademlia.Alpha; i++ {
				var contact = shortList[i]
				if findValue {
					shortList, probed, value = kademlia.sendFindValueRPCs(network, &contact, target, shortList, probed)
					if len(value) != 0 {
						return nil, value
					}
				} else {
					shortList, probed = kademlia.sendFindNodeRPCs(network, contact, &targetID, shortList, probed)
				}
			}

		} else { // if more than k nodes has been successfully probed, send RPCs to k closest (not yet probed) //todo: if closest unchanged, outside of while
			//if
			for i := 0; i < kademlia.K; i++ {
				var contact = shortList[i]
				if findValue {
					shortList, probed, value = kademlia.sendFindValueRPCs(network, &contact, target, shortList, probed)
					if len(value) != 0 {
						return nil, value
					}
				} else {
					shortList, probed = kademlia.sendFindNodeRPCs(network, contact, &targetID, shortList, probed)
				}
			}
		}
	}
	//todo: return shortlist
}*/

func (kademlia *Kademlia)sendFindNodeRPCs(network Network, contact Contact, id *KademliaID, shortList []Contact, probed []Contact)(newShortList []Contact, newProbed []Contact) {
	var received = FindNodeRPC(network, contact.Address, id.String(), kademlia.Me)


	//fmt.Printf("RECEIVED, nodeLookup, sendFindNodeRPCs: " + printContacts(received) + "\n")
	newProbed = append(probed, contact)
	newShortList = updateShortList(received, id, shortList, probed)
	//fmt.Printf("NEWSHORTLIST, nodeLookup, sendFindNodeRPCs: " + printContacts(newShortList) + "\n")
	//fmt.Printf("NEWPROBED, nodeLookup, sendFindNodeRPCs: " + printContacts(newProbed) + "\n")
	//return newShortList, newProbed

	return received, newProbed
}

func (kademlia *Kademlia)sendFindValueRPCs(network Network, contact *Contact, hash string, shortList []Contact, probed []Contact)(newShortList []Contact, newProbed []Contact, value []byte) {
	var id = NewKademliaID(hash)
	value, received  := FindValueRPC(network, contact.Address, hash, kademlia.Me)

	if len(value) == 0 {
		newProbed = append(probed, *contact)
		newShortList = updateShortList(received, id, shortList, probed)
	}
	return newShortList, newProbed, value
}

func updateShortList(contacts []Contact, id *KademliaID, shortList []Contact, probed []Contact)(newShortList []Contact) {
	for _, contact := range contacts {
		var alreadyProbed = false
		for _, used := range probed {
			if contact.ID.Equals(used.ID) {
				alreadyProbed = true
			}
		}
		if !alreadyProbed {
			newShortList = append(shortList, contact)
		}
	}
	newShortList = sortContacts(id, newShortList)
	return newShortList
}

func sortContacts(id *KademliaID, unsorted []Contact)(sorted []Contact) {
	candidates := ShortList{id, unsorted}
	candidates.Sort()
	sorted = candidates.Contacts
	return sorted
}
