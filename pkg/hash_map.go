package pkg

import "hash-map/pkg/hashing"

type HashMap struct {
	buffer [10]*ArrayList
}

func NewHashMap() *HashMap {
	return &HashMap{}
}

func (hp *HashMap) Add(value string) {
	hash := hashing.CalculateHash64([]byte(value))

	index := hp.getIndexOfMap(hash)
	for i, current := range hp.buffer {
		if uint64(i) == index {
			if current == nil {
				hp.buffer[index] = &ArrayList{
					Actual: value,
				}
			} else {
				current.Append(value)
			}
		}
	}
}

func (hp *HashMap) Get(value string) *string {
	hash := hashing.CalculateHash64([]byte(value))

	index := hp.getIndexOfMap(hash)
	linkedList := hp.buffer[index]
	if linkedList == nil {
		return nil
	}

	return linkedList.Get(value)
}

func (hp *HashMap) getIndexOfMap(hash uint64) uint64 {
	return hash % uint64(len(hp.buffer))
}
