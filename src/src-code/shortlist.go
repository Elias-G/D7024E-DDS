package src

import "sort"

type ShortList struct {
	Target   *KademliaID
	Contacts []Contact
}

// Sort the Contacts in ContactCandidates
func (candidates *ShortList) Sort() {
	sort.Sort(candidates)
}

// Len returns the length of the ContactCandidates
func (candidates *ShortList) Len() int {
	return len(candidates.Contacts)
}

// Swap the position of the Contacts at i and j
// WARNING does not check if either i or j is within range
func (candidates *ShortList) Swap(i, j int) {
	candidates.Contacts[i], candidates.Contacts[j] = candidates.Contacts[j], candidates.Contacts[i]
}

// Less returns true if the Contact at index i is closer to target than
// the Contact at index j
func (candidates *ShortList) Less(i, j int) bool {
	return candidates.Contacts[i].ID.CalcDistance(candidates.Target).Less(candidates.Contacts[j].ID.CalcDistance(candidates.Target))
}
