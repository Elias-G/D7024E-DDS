package src

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestHashValue(t *testing.T) {
	value,_ := hex.DecodeString("TestHash")
	hash := HashValue(value)
	if len(hash) == 0 {
		t.Errorf("HashValue failed, length 0")
	}

	var i interface {} = hash

	_,ok := i.(string)
	if !ok {
		t.Errorf("HashValue failed, did not return string")
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