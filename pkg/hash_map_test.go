package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap_WhenExistingValue(f *testing.T) {
	hashMap := NewHashMap()
	hashMap.Add("oi")

	assert.NotNil(f, hashMap.Get("oi"))
}

func TestHashMap_NonExistingValue(f *testing.T) {
	hashMap := NewHashMap()

	assert.Nil(f, hashMap.Get("oi"))
}
