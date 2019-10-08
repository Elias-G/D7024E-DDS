package src

import (
	"crypto/sha1"
	"encoding/hex"
)

type Storage struct{}

func InitTable() (Table map[string][]byte) {
	Table = make(map[string][]byte)
	return Table
}

func HashValue(value []byte) (key string) {
	hash := sha1.New()
	v := hash.Sum(value)
	key = hex.EncodeToString(v)
	return key
}
