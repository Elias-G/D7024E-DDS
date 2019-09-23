package src

import "crypto/sha1"

type Storage struct{}

func InitTable() (Table map[string][]byte) {
	Table = make(map[string][]byte)
	return Table
}

func HashValue(value []byte) (key string) {
	hash := sha1.New()
	hash.Write(value)
	key = string(hash.Sum(nil))
	return key
}
