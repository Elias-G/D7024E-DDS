package src

import (
	"crypto/sha1"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestHashValue(t *testing.T) {
	value,_ := hex.DecodeString("TestHash")
	hash := sha1.New()
	hash.Write(value)
	want := string(hash.Sum(nil))
	got := HashValue(value)
	if want != got {
		t.Errorf("HashValue() = %v, want %v", got, want)
	}
}

func TestInitTable(t *testing.T) {
	tests := []struct {
		name      string
		wantTable map[string][]byte
	}{
		{"TestTable", make(map[string][]byte)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTable := InitTable(); !reflect.DeepEqual(gotTable, tt.wantTable) {
				t.Errorf("InitTable() = %v, want %v", gotTable, tt.wantTable)
			}
		})
	}
}