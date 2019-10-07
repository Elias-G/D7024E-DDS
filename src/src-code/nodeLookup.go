package src

// Sends out alpha RPCs for FindNode/FindValue and returns k closest contacts or value if found
// TODO: Parallel requests, Timing?
func (kademlia *Kademlia) NodeLookup(network Network, target *Contact, hash string)(contacts []Contact, value []byte) {
	var table = kademlia.RoutingTable
	var alpha = kademlia.Alpha
	var k = kademlia.K
	var shortList []Contact
	var probed []Contact // keep track of contacts that have already been probed
	var noVal []byte // If no value was found or searched for, return empty byte array
	var targetID *KademliaID

	shortList, targetID = initShortlist(table, target, hash, alpha)
	var closestNode = shortList[0].ID // current closest node to target

	// TODO: Parallel?
	for i := 0; i < alpha; i++ {
		var contact = shortList[i]
		if target == nil {
			shortList, probed, value = kademlia.sendFindValueRPCs(network, &contact, hash, shortList, probed)
			if len(value) != 0 {
				return nil, value
			}
		} else {
			shortList, probed = kademlia.sendFindNodeRPCs(network, contact, target.ID, shortList, probed)
		}
	}

	// While recent responses are closer than closestNode, send new RPCs
	for shortList[0].ID.CalcDistance(targetID).Less(closestNode.CalcDistance(targetID)) {
		closestNode = shortList[0].ID
		// if less than k nodes has been successfully probed
		if len(probed) < k {
			for i := 0; i < alpha; i++ {
				var contact = shortList[i]
				if contacts == nil {
					shortList, probed, value = kademlia.sendFindValueRPCs(network, &contact, hash, shortList, probed)
					if len(value) != 0 {
						return nil, value
					}
				} else {
					shortList, probed = kademlia.sendFindNodeRPCs(network, contact, target.ID, shortList, probed)
				}
			}
		} else { // if more than k nodes has been successfully probed, send RPCs to k closest (not yet probed)
			for i := 0; i < k; i++ {
				var contact = shortList[i]
				if contacts == nil {
					shortList, probed, value = kademlia.sendFindValueRPCs(network, &contact, hash, shortList, probed)
					if len(value) != 0 {
						return nil, value
					}
				} else {
					shortList, probed = kademlia.sendFindNodeRPCs(network, contact, target.ID, shortList, probed)
				}
			}
		}
	}

	return contacts, noVal
}

func initShortlist(table RoutingTable, target *Contact, hash string, alpha int)([]Contact, *KademliaID) {
	var shortList []Contact
	var id *KademliaID
	if target != nil {
		id = target.ID
		shortList = table.FindClosestContacts(id, alpha)
	} else {
		id = NewKademliaID(hash)
		shortList = table.FindClosestContacts(id, alpha)
	}
	id = shortList[0].ID
	return shortList, id
}

func (kademlia *Kademlia)sendFindNodeRPCs(network Network, contact Contact, id *KademliaID, shortList []Contact, probed []Contact)(newShortList []Contact, newProbed []Contact) {
	var received = FindNodeRPC(network, contact.Address, id.String(), kademlia.Me)
	newProbed = append(probed, contact)
	newShortList = updateShortList(received, id, shortList, probed)
	return newShortList, newProbed
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

// Sort shortList
func sortContacts(id *KademliaID, unsorted []Contact)(sorted []Contact) {
	candidates := ShortList{id, unsorted}
	candidates.Sort()
	sorted = candidates.Contacts
	return sorted
}
