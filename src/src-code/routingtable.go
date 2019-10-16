package src

import "fmt"
import "strconv"

var bucketSize int

// RoutingTable definition
// keeps a refrence contact of me and an array of buckets
type RoutingTable struct {
	me      Contact
	buckets [IDLength * 8]*bucket
}

// NewRoutingTable returns a new instance of a RoutingTable
func NewRoutingTable(me Contact, k int) *RoutingTable {
	routingTable := &RoutingTable{}
	for i := 0; i < IDLength*8; i++ {
		routingTable.buckets[i] = newBucket()
	}
	routingTable.me = me
	bucketSize = k
	return routingTable
}

func (routingTable *RoutingTable) UpdateRoutingTable(sender Contact) bool {
	success := routingTable.AddContact(sender)
	if success {
		return true
	}
	return false

}

// AddContact add a new contact to the correct Bucket
func (routingTable *RoutingTable) AddContact(contact Contact) bool {
	bucketIndex := routingTable.getBucketIndex(contact.ID)
	fmt.Print("BucketIndex that contact is added to: " + strconv.Itoa(bucketIndex))
	bucket := routingTable.buckets[bucketIndex]
	if !bucket.full() {
		bucket.AddContact(contact)
		return true
	}
	return false
}

// RemoveContact remove a contact from the correct Bucket
func (routingTable *RoutingTable) RemoveContact(contact Contact) {
	bucketIndex := routingTable.getBucketIndex(contact.ID)
	//fmt.Print("BucketIndex that contact is to be removed from: " + strconv.Itoa(bucketIndex))
	bucket := routingTable.buckets[bucketIndex]
	//fmt.Print("Contact to remove: " + contact.String() + "\n")
	//fmt.Print("Bucket before remove: " + bucket.String() + "\n")
	//fmt.Print("Routingtable before Remove: " + routingTable.String() + "\n")
	bucket.RemoveContact(contact)
	//fmt.Print("Bucket after remove: " + bucket.String() + "\n")
	//fmt.Print("Routingtable after Remove: " + routingTable.String() + "\n")
}

// FindClosestContacts finds the count closest Contacts to the target in the RoutingTable
func (routingTable *RoutingTable) FindClosestContacts(target *KademliaID, count int) []Contact {
	var candidates ContactCandidates
	bucketIndex := routingTable.getBucketIndex(target)
	bucket := routingTable.buckets[bucketIndex]

	candidates.Append(bucket.GetContactAndCalcDistance(target))

	for i := 1; (bucketIndex-i >= 0 || bucketIndex+i < IDLength*8) && candidates.Len() < count; i++ {
		if bucketIndex-i >= 0 {
			bucket = routingTable.buckets[bucketIndex-i]
			candidates.Append(bucket.GetContactAndCalcDistance(target))
		}
		if bucketIndex+i < IDLength*8 {
			bucket = routingTable.buckets[bucketIndex+i]
			candidates.Append(bucket.GetContactAndCalcDistance(target))
		}
	}

	candidates.Sort()

	if count > candidates.Len() {
		count = candidates.Len()
	}

	return candidates.GetContacts(count)
}

// getBucketIndex get the correct Bucket index for the KademliaID
func (routingTable *RoutingTable) getBucketIndex(id *KademliaID) int {
	distance := id.CalcDistance(routingTable.me.ID)
	for i := 0; i < IDLength; i++ {
		for j := 0; j < 8; j++ {
			if (distance[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}

	return IDLength*8 - 1
}

func (routingTable *RoutingTable) String() string {
	contacts := routingTable.FindClosestContacts(routingTable.me.ID, 20)
	answer := "\n"
	for _, contact := range contacts {
		answer += "Address: " + contact.Address + "\n"
	}
	return answer
}
