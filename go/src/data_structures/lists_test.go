package data_structures_test

import (
	"github.com/99_problems/go/src/data_structures"
	"github.com/99_problems/go/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	testList(t, data_structures.NewArrayList())
	testList(t, data_structures.NewLinkedList())
}

func testList(t *testing.T, list data_structures.List) {
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, true, list.IsEmpty())

	list.Add(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, false, list.IsEmpty())

	list.Remove(0)
	assert.Equal(t, 0, list.Size())
	assert.Equal(t, true, list.IsEmpty())
}

func TestListOverflow(t *testing.T) {
	testOverflow(t, data_structures.NewArrayList())
	testOverflow(t, data_structures.NewLinkedList())
}

func testOverflow(t *testing.T, list data_structures.List) {
	for i := 0; i < 100; i++ {
		list.Add(i)
	}

	for i := 0; i < 100; i++ {
		assert.Equal(t, list.Get(i), i)
	}

	list.Remove(0)
	assert.Equal(t, 99, list.Size())

	for i := 0; i < 99; i++ {
		assert.Equal(t, i+1, list.Get(i))
	}
}

func TestListIterator(t *testing.T) {
	testIterator(t, data_structures.NewLinkedList())
	testIterator(t, data_structures.NewArrayList())
}

func testIterator(t *testing.T, list data_structures.List) {
	list.Add(1)
	list.Add(2)
	list.Add(3)

	iterator := list.Iterator()
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 1, iterator.Next())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 2, iterator.Next())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, 3, iterator.Next())
	assert.False(t, iterator.HasNext())
}

func TestEmptyIterator(t *testing.T) {
	testEmptyIterator(t, data_structures.NewLinkedList())
	testEmptyIterator(t, data_structures.NewArrayList())
}

func testEmptyIterator(t *testing.T, list data_structures.List) {
	iterator := list.Iterator()
	assert.False(t, iterator.HasNext())
}

func TestFilter(t *testing.T) {
	testFilter(t, data_structures.NewLinkedList())
	testFilter(t, data_structures.NewArrayList())
}

func testFilter(t *testing.T, list data_structures.List) {
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	filtered := list.Filter(func(value utils.Any) bool {
		if value.(int) >= 2 {
			return true
		}
		return false
	})

	assert.Equal(t, 3, filtered.Size())

	assert.Equal(t, 2, filtered.Get(0))
	assert.Equal(t, 3, filtered.Get(1))
	assert.Equal(t, 4, filtered.Get(2))
}
