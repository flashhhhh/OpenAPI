package bloomfilter

import "hash/fnv"

type BloomFilter struct {
	hashCount int
	size int
	bitSet    []bool
}

func NewBloomFilter(size, hashCount int) *BloomFilter {
	return &BloomFilter{
		hashCount: hashCount,
		size: size,
		bitSet: make([]bool, size),
	}
}

func (bf *BloomFilter) hash(value string, seed int) int {
	h := fnv.New64a()
	h.Write([]byte(value))
	h.Write([]byte{byte(seed)})
	return int(h.Sum64() % uint64(bf.size))
}

func (bf *BloomFilter) Add(value string) {
	for i := 0; i < bf.hashCount; i++ {
		index := bf.hash(value, i)
		bf.bitSet[index] = true
	}
}

func (bf *BloomFilter) Contains(value string) bool {
	for i := 0; i < bf.hashCount; i++ {
		index := bf.hash(value, i)
		if !bf.bitSet[index] {
			return false
		}
	}
	return true
}