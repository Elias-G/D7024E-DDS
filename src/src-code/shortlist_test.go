package src

import (
	"reflect"
	"testing"
)

func TestShortList_Append(t *testing.T) {
	type fields struct {
		Target   *KademliaID
		Contacts []Contact
	}
	type args struct {
		contacts []Contact
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ShortList{
				Target:   tt.fields.Target,
				Contacts: tt.fields.Contacts,
			}
		})
	}
}

func TestShortList_GetContacts(t *testing.T) {
	type fields struct {
		Target   *KademliaID
		Contacts []Contact
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Contact
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ShortList{
				Target:   tt.fields.Target,
				Contacts: tt.fields.Contacts,
			}
			if got := candidates.GetContacts(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortList_Len(t *testing.T) {
	type fields struct {
		Target   *KademliaID
		Contacts []Contact
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ShortList{
				Target:   tt.fields.Target,
				Contacts: tt.fields.Contacts,
			}
			if got := candidates.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortList_Less(t *testing.T) {
	type fields struct {
		Target   *KademliaID
		Contacts []Contact
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ShortList{
				Target:   tt.fields.Target,
				Contacts: tt.fields.Contacts,
			}
			if got := candidates.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortList_Sort(t *testing.T) {
	type fields struct {
		Target   *KademliaID
		Contacts []Contact
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ShortList{
				Target:   tt.fields.Target,
				Contacts: tt.fields.Contacts,
			}
		})
	}
}

func TestShortList_Swap(t *testing.T) {
	type fields struct {
		Target   *KademliaID
		Contacts []Contact
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := &ShortList{
				Target:   tt.fields.Target,
				Contacts: tt.fields.Contacts,
			}
		})
	}
}