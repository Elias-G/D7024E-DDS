package src

// Sends out alpha RPCs for FindNode/FindValue and returns k closest contacts or value if found
func (kademlia *Kademlia) NodeLookup(network Network, target string, findValue bool) (contacts []Contact, value []byte) {
	var table = kademlia.RoutingTable
	var alpha = kademlia.Alpha
	var shortList []Contact
	var probed []Contact // keep track of contacts that have already been probed
	var targetID *KademliaID

	targetID = NewKademliaID(target)
	shortList = table.FindClosestContacts(targetID, alpha)
	for i := 0; i < len(shortList) && i < alpha; i++ {
		var contact = shortList[i]
		if findValue {
			shortList, probed, value = kademlia.sendFindValueRPCs(network, &contact, target, shortList, probed)
			if len(value) != 0 {
				return nil, value
			}
		} else {
			shortList, probed = kademlia.sendFindNodeRPCs(network, contact, targetID, shortList, probed)
		}
	}

	contacts, value = kademlia.iterativeLookup(network, shortList, probed, target, *targetID, shortList[0].ID, findValue, value)
	if network.Node.K < len(contacts) {
		return contacts[0:network.Node.K], value
	} else {
		return contacts, value
	}
}

func (kademlia *Kademlia) iterativeLookup(network Network, shortList []Contact, probed []Contact, target string, targetID KademliaID, closestNode *KademliaID, findValue bool, value []byte) ([]Contact, []byte) {
	// While recent responses are closer than closestNode, and less than k nodes has been successfully probed, send new RPCs
	firsttime := true
	for (firsttime || shortList[0].ID.CalcDistance(&targetID).Less(closestNode.CalcDistance(&targetID))) && (len(probed) < kademlia.K) {
		firsttime = false
		closestNode = shortList[0].ID
		for i := 0; i < len(shortList) && i < kademlia.Alpha; i++ {
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
	// if more than k nodes has been successfully probed, send RPCs to k closest (not yet probed)
	if len(probed) >= kademlia.K {
		for i := 0; i < len(shortList) && i < kademlia.K && i < len(shortList); i++ {
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

	return shortList, value
}

func (kademlia *Kademlia) sendFindNodeRPCs(network Network, contact Contact, id *KademliaID, shortList []Contact, probed []Contact) (newShortList []Contact, newProbed []Contact) {
	var received = FindNodeRPC(network, contact.Address, id.String(), kademlia.Me)
	newProbed = append(probed, contact)
	newShortList = updateShortList(received, id, shortList, probed)
	return newShortList, newProbed
}

func (kademlia *Kademlia) sendFindValueRPCs(network Network, contact *Contact, hash string, shortList []Contact, probed []Contact) (newShortList []Contact, newProbed []Contact, value []byte) {
	var id = NewKademliaID(hash)
	value, received := FindValueRPC(network, contact.Address, hash, kademlia.Me)

	if len(value) == 0 {
		newProbed = append(probed, *contact)
		newShortList = updateShortList(received, id, shortList, probed)
	}
	return newShortList, newProbed, value
}

func updateShortList(received []Contact, id *KademliaID, shortList []Contact, probed []Contact) (newShortList []Contact) {
	newShortList = shortList
	for _, contact := range received {
		alreadyProbed := inList(contact, probed)
		alreadyInShortList := inList(contact, newShortList)
		if alreadyProbed == false && alreadyInShortList == false {
			newShortList = append(newShortList, contact)
		}
	}
	newShortList = sortContacts(id, newShortList)
	return newShortList
}

func inList(contact Contact, list []Contact) bool {
	for _, c := range list {
		if c.ID.Equals(contact.ID) {
			return true
		}
	}
	return false
}

func sortContacts(id *KademliaID, unsorted []Contact) (sorted []Contact) {
	candidates := ShortList{id, unsorted}
	candidates.Sort()
	sorted = candidates.Contacts
	return sorted
}
