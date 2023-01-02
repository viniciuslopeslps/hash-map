package hashing

import (
	"hash/fnv"
)

func CalculateHash64(value []byte) uint64 {
	hash := fnv.New64a()

	hash.Write(value)
	return hash.Sum64()
}
