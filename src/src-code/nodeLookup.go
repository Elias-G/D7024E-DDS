package src

// Sends out alpha RPCs for FindNode/FindValue and returns k closest contacts or value if found
// TODO: Parallel requests, support for findValue, keep track of nodes already probed? Timing? How to return value or contacts?
func (kademlia *Kademlia) NodeLookup(network Network, target *Contact, hash string, findValue bool)(contacts []Contact, value string) {
	var table = kademlia.RoutingTable
	var alpha = kademlia.Alpha
	var k = kademlia.K
	var id = target.ID
	var shortList = table.FindClosestContacts(id, alpha) // initiating shortList
	var closestNode = shortList[0].ID // current closest node to target
	var probed []Contact // keep track of contacts that have already been probed

	// TODO: Parallel?
	for i := 0; i < alpha; i++ {
		var contact = shortList[i]
		if findValue {
			shortList, probed, value = kademlia.sendFindValueRPCs(network, contact, id, shortList, probed, findValue)
			if value != "" {
				return nil, value
			}
		} else {
			shortList, probed = kademlia.sendFindNodeRPCs(network, contact, id, shortList, probed)
		}
	}

	// While recent responses are closer than closestNode, send new RPCs
	for shortList[0].ID.CalcDistance(id).Less(closestNode.CalcDistance(id)) {
		closestNode = shortList[0].ID
		// if less than k nodes has been successfully probed
		if len(probed) < k {
			for i := 0; i < alpha; i++ {
				var contact = shortList[i]
				if findValue {
					shortList, probed, value = kademlia.sendFindValueRPCs(network, contact, id, shortList, probed, findValue)
					if value != "" {
						return nil, value
					}
				} else {
					shortList, probed = kademlia.sendFindNodeRPCs(network, contact, id, shortList, probed)
				}
			}
		} else { // if more than k nodes has been successfully probed, send RPCs to k closest (not yet probed)
			for i := 0; i < k; i++ {
				var contact = shortList[i]
				if findValue {
					shortList, probed, value = kademlia.sendFindValueRPCs(network, contact, id, shortList, probed, findValue)
					if value != "" {
						return nil, value
					}
				} else {
					shortList, probed = kademlia.sendFindNodeRPCs(network, contact, id, shortList, probed)
				}
			}
		}
	}

	return contacts, ""
}

func (kademlia *Kademlia)sendFindNodeRPCs(network Network, contact Contact, id *KademliaID, shortList []Contact, probed []Contact)(newShortList []Contact, newProbed []Contact) {
	var received = FindNodeRPC(network, contact.Address, id.String(), kademlia.Me)
	newProbed = append(probed, contact)
	newShortList = updateShortList(received, id, shortList, probed)
	return newShortList, newProbed
}

func (kademlia *Kademlia)sendFindValueRPCs(network Network, contact Contact, id *KademliaID, shortList []Contact, probed []Contact, findValue bool)(newShortList []Contact, newProbed []Contact, value string) {
	value = ""
	received := FindValueRPC(network, contact.Address, id.String(), kademlia.Me)
	var i interface {} = received
	_, foundValue := i.(string)
	if !foundValue {
		newProbed = append(probed, contact)
		//newShortList = updateShortList(received, id, shortList, probed) //todo: findvalue rpc returns a value not contacts??
	} else {
		value = 	"RANDOM VALUE"//received //todo ??
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

// Sort shortList
func sortContacts(id *KademliaID, unsorted []Contact)(sorted []Contact) {
	candidates := ShortList{id, unsorted}
	candidates.Sort()
	sorted = candidates.Contacts
	return sorted
}
