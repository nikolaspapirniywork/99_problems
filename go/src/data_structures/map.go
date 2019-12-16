package data_structures

import (
	"github.com/99_problems/go/src/utils"
	"math"
)

type Map interface {
	Put(key utils.Any, value utils.Any)
	Get(key utils.Any) utils.Any
	Remove(key utils.Any) utils.Any
	Size() int
	Capacity() int
	IsEmpty() bool
	BucketDistribution() List
	Ratio() float64
}

const mapInitialCapacity = 10
const lowerThreshold = 0.3
const upperThreshold = 0.7

type MapEntry struct {
	Key   utils.Any
	Value utils.Any
}

func CopyMapEntry(entry MapEntry) *MapEntry {
	return &MapEntry{Key: entry.Key, Value: entry.Value}
}

type HashMap struct {
	buckets []*LinkedList
	size    int
}

func NewHashMap() *HashMap {
	buckets := initBuckets(mapInitialCapacity)
	return &HashMap{buckets, 0}
}

func initBuckets(size int) []*LinkedList {
	buckets := make([]*LinkedList, size)
	for i, _ := range buckets {
		buckets[i] = NewLinkedList()
	}
	return buckets
}

func (m HashMap) Get(key utils.Any) utils.Any {
	bucket := m.getBucket(key)
	if bucket.IsEmpty() {
		return nil
	}

	iterator := bucket.Iterator()

	for iterator.HasNext() {
		entry := iterator.Next().(*MapEntry)
		if entry.Key == key {
			return entry.Value
		}
	}

	return nil
}

func (m *HashMap) Remove(key utils.Any) utils.Any {
	bucket := m.getBucket(key)
	if bucket.IsEmpty() {
		return nil
	}

	foundEntry := bucket.find(func(value utils.Any) bool {
		if value.(*MapEntry).Key == key {
			return true
		}
		return false
	})

	if foundEntry == nil {
		return nil
	}

	value := foundEntry.Value
	bucket.remove(foundEntry)
	m.size -= 1

	m.ensureTresholdMet()

	return value
}

func (m *HashMap) Put(key utils.Any, value utils.Any) {
	bucket := m.getBucket(key)
	bucket.Add(&MapEntry{key, value})
	m.size += 1
	m.ensureTresholdMet()
}

func (m *HashMap) ensureTresholdMet() {
	ratio := m.Ratio()

	if ratio > upperThreshold {
		m.increaseSizeAndRehash()
		return
	}

	if m.Size() > mapInitialCapacity && ratio < lowerThreshold {
		m.decreaseSizeAndRehash()
	}
}

func (m *HashMap) Ratio() float64 {
	return float64(m.Size()) / float64(m.Capacity())
}

func (m *HashMap) increaseSizeAndRehash() {
	newBuckets := initBuckets(m.Capacity() * 2)
	rehash(m.buckets, newBuckets)
	m.buckets = newBuckets
}

func (m *HashMap) decreaseSizeAndRehash() {
	newSize := m.Capacity() / 2
	if newSize < mapInitialCapacity {
		newSize = mapInitialCapacity
	}

	newBuckets := initBuckets(newSize)
	rehash(m.buckets, newBuckets)
	m.buckets = newBuckets
}

func rehash(oldBuckets []*LinkedList, newBuckets []*LinkedList) {
	for _, bucket := range oldBuckets {
		iterator := bucket.Iterator()
		for iterator.HasNext() {
			next := iterator.Next()
			entry := next.(*MapEntry)
			newIndex := getBucketIndex(entry.Key, len(newBuckets))
			newBuckets[newIndex].Add(CopyMapEntry(*entry))
		}
	}
}

func (m HashMap) Size() int {
	return m.size
}

func (m HashMap) Capacity() int {
	return len(m.buckets)
}

func (m HashMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m HashMap) getBucket(key utils.Any) *LinkedList {
	bucketIndex := getBucketIndex(key, m.Capacity())
	return m.buckets[bucketIndex]
}

func getBucketIndex(key utils.Any, capacity int) int {
	keyHash := utils.HashObject(key)
	return int(math.Abs(float64(keyHash % capacity)))
}

func (m HashMap) BucketDistribution() List {
	result := NewArrayList()
	for _, bucket := range m.buckets {
		result.Add(bucket.Size())
	}
	return result
}
