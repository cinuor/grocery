package bloomfilter

import (
	"hash"
	"hash/crc64"
	"hash/fnv"
	"log"
	"math"
)

// "math"

type BloomFilter struct {
	Bytes []byte
	Hashs []hash.Hash
}

func (f *BloomFilter) Push() error {
	
}

