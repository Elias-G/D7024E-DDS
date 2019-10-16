package src

import (
	"fmt"
	"testing"
)

//Test to remove contact that does not exist in bucket
func TestBucket_remove_nonexisting(t *testing.T) {
	bucket := newBucket()
	fmt.Println(bucket.Len())
	contact1 := NewContact(NewKademliaID("1111111100000000000000000000000000000001"), "localhost:8002")
	contact2 := NewContact(NewKademliaID("1111111100000000000000000000000000000002"), "localhost:8002")
	contact3 := NewContact(NewKademliaID("1111111100000000000000000000000000000003"), "localhost:8002")
	contact4 := NewContact(NewKademliaID("1111111100000000000000000000000000000004"), "localhost:8002")

	bucket.AddContact(contact1)
	bucket.AddContact(contact2)
	bucket.AddContact(contact3)

	//fmt.Println(bucket.Len())
	bucket.RemoveContact(contact4)
	//fmt.Println(bucket.String())
	got := bucket.String()
	wantedbucket := newBucket()
	wantedbucket.AddContact(contact1)
	wantedbucket.AddContact(contact2)
	wantedbucket.AddContact(contact3)
	if got != wantedbucket.String() {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", got, wantedbucket)
	}
}

//Test remove first contact added to the bucket
func TestBucket_remove_firstadded(t *testing.T) {
	bucket := newBucket()
	//fmt.Println(bucket.Len())
	contact1 := NewContact(NewKademliaID("733e23fb94bcff1a249cb95ab634cbd6c2beeba2"), "localhost:8002")
	contact2 := NewContact(NewKademliaID("1111111100000000000000000000000000000002"), "localhost:8002")
	contact3 := NewContact(NewKademliaID("1111111100000000000000000000000000000003"), "localhost:8002")

	bucket.AddContact(contact1)
	bucket.AddContact(contact2)
	bucket.AddContact(contact3)

	bucket.RemoveContact(contact1)
	//fmt.Println(bucket.String())

	wantedbucket := newBucket()
	wantedbucket.AddContact(contact2)
	wantedbucket.AddContact(contact3)
	if bucket.String() != wantedbucket.String() {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", bucket, wantedbucket)
	}
}

func TestBucket_remove_middle(t *testing.T) {
	bucket := newBucket()
	//fmt.Println(bucket.Len())
	contact1 := NewContact(NewKademliaID("733e23fb94bcff1a249cb95ab634cbd6c2beeba2"), "localhost:8002")
	contact2 := NewContact(NewKademliaID("1111111100000000000000000000000000000002"), "localhost:8002")
	contact3 := NewContact(NewKademliaID("1111111100000000000000000000000000000003"), "localhost:8002")

	bucket.AddContact(contact1)
	bucket.AddContact(contact2)
	bucket.AddContact(contact3)

	bucket.RemoveContact(contact2)
	//fmt.Println(bucket.String())

	wantedbucket := newBucket()
	wantedbucket.AddContact(contact1)
	wantedbucket.AddContact(contact3)
	if bucket.String() != wantedbucket.String() {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", bucket, wantedbucket)
	}

}

//Test remove last contact added to the bucket
func TestBucket_bucket_full(t *testing.T) {
	bucket := newBucket()
	//fmt.Println(bucket.Len())
	n := int64(1095216660481)
	for i := 0; i < 20; i++ {
		n++
		s := fmt.Sprintf("%b", n)
		//fmt.Printf("%b\n", n)
		contact1 := NewContact(NewKademliaID(s), "localhost:8002")
		bucket.AddContact(contact1)
	}
	got := bucket.full()
	want := true
	if got != want {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", got, want)
	}
}

func TestBucket_getHead(t *testing.T) {
	bucket := newBucket()
	contact1 := NewContact(NewKademliaID("1111111100000000000000000000000000000001"), "localhost:8001")
	contact2 := NewContact(NewKademliaID("1111111100000000000000000000000000000002"), "localhost:8002")
	bucket.AddContact(contact1)
	bucket.AddContact(contact2)
	got := bucket.getHead().ID.String()
	want := contact2.ID.String()
	if got != want {
		t.Errorf("KademliaID_Equals(id) = %v, want %v", got, want)
	}
}
