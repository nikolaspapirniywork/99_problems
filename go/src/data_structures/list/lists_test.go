package list_test

import (
	"github.com/99_problems/go/src/data_structures"
	"github.com/99_problems/go/src/data_structures/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	testList(t, list.NewArrayList())
	testList(t, list.NewLinkedList())
}

func testList(t *testing.T, list list.List) {
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
	testOverflow(t, list.NewArrayList())
	testOverflow(t, list.NewLinkedList())
}

func testOverflow(t *testing.T, list list.List) {
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
	testIterator(t, list.NewLinkedList())
	testIterator(t, list.NewArrayList())
}

func testIterator(t *testing.T, list list.List) {
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
	testEmptyIterator(t, list.NewLinkedList())
	testEmptyIterator(t, list.NewArrayList())
}

func testEmptyIterator(t *testing.T, list list.List) {
	iterator := list.Iterator()
	assert.False(t, iterator.HasNext())
}

func TestFilter(t *testing.T) {
	testFilter(t, list.NewLinkedList())
	testFilter(t, list.NewArrayList())
}

func testFilter(t *testing.T, list list.List) {
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	filtered := list.Filter(func(value data_structures.Any) bool {
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
