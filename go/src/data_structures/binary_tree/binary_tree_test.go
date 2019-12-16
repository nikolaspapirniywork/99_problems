package binary_tree_test

import (
	"github.com/99_problems/go/src/data_structures"
	"github.com/99_problems/go/src/data_structures/binary_tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ComparableInt int

func (i ComparableInt) LessThan(value data_structures.Any) bool {
	return i < (value.(ComparableInt))
}

func (i ComparableInt) GreaterThan(value data_structures.Any) bool {
	return i > (value.(ComparableInt))
}

func TestComparableInt(t *testing.T) {
	assert.True(t, ComparableInt(2).LessThan(ComparableInt(3)))
	assert.False(t, ComparableInt(2).LessThan(ComparableInt(2)))
	assert.False(t, ComparableInt(3).LessThan(ComparableInt(2)))

	assert.False(t, ComparableInt(2).GreaterThan(ComparableInt(3)))
	assert.False(t, ComparableInt(2).GreaterThan(ComparableInt(2)))
	assert.True(t, ComparableInt(3).GreaterThan(ComparableInt(2)))
}

func TestBinaryTreeEmpty(t *testing.T) {
	tree := binary_tree.New()
	assert.Equal(t, 0, tree.Size())
	assert.True(t, tree.IsEmpty())
}

func TestBinaryTreeAdd(t *testing.T) {
	tree := binary_tree.New()

	tree.Add(ComparableInt(2))
	tree.Add(ComparableInt(6))
	assert.Equal(t, 2, tree.Size())
}

func TestBinaryTreeRemoveHead(t *testing.T) {
	tree := binary_tree.New()
	tree.Add(ComparableInt(2))
	tree.Remove(ComparableInt(2))

	assert.Equal(t, 0, tree.Size())
	assert.True(t, tree.IsEmpty())
}

func TestBinaryTreeRemoveLeft(t *testing.T) {
	tree := binary_tree.New()
	tree.Add(ComparableInt(5))
	tree.Add(ComparableInt(2))
	tree.Remove(ComparableInt(2))

	assert.Equal(t, 1, tree.Size())
	assert.False(t, tree.IsEmpty())
}

func TestBinaryTreeRemoveLeftMultiple(t *testing.T) {
	tree := binary_tree.New()

	tree.Add(ComparableInt(5))
	tree.Add(ComparableInt(3))
	tree.Add(ComparableInt(2))
	tree.Add(ComparableInt(1))

	tree.Remove(ComparableInt(1))
	tree.Remove(ComparableInt(2))
	tree.Remove(ComparableInt(3))

	assert.Equal(t, 1, tree.Size())
	assert.False(t, tree.IsEmpty())
}

func TestBinaryTreeRemoveRight(t *testing.T) {
	tree := binary_tree.New()
	tree.Add(ComparableInt(5))
	tree.Add(ComparableInt(10))
	tree.Remove(ComparableInt(10))

	assert.Equal(t, 1, tree.Size())
	assert.False(t, tree.IsEmpty())
}

func TestBinaryTreeRemoveRightMultiple(t *testing.T) {
	tree := binary_tree.New()

	tree.Add(ComparableInt(5))
	tree.Add(ComparableInt(10))
	tree.Add(ComparableInt(20))
	tree.Add(ComparableInt(30))

	tree.Remove(ComparableInt(10))
	tree.Remove(ComparableInt(20))
	tree.Remove(ComparableInt(30))

	assert.Equal(t, 1, tree.Size())
	assert.False(t, tree.IsEmpty())
}

func TestBinaryTreeRemoveBothSides(t *testing.T) {
	tree := binary_tree.New()

	tree.Add(ComparableInt(5))
	tree.Add(ComparableInt(2))
	tree.Add(ComparableInt(10))

	tree.Remove(ComparableInt(2))
	tree.Remove(ComparableInt(10))
	tree.Remove(ComparableInt(5))

	assert.Equal(t, 0, tree.Size())
	assert.True(t, tree.IsEmpty())
}

func TestBinaryTreeRemoveManyValues(t *testing.T) {
	tree := binary_tree.New()

	for i := 0; i < 100; i++ {
		tree.Add(ComparableInt(i))
	}
	assert.Equal(t, 100, tree.Size())

	for i := 0; i < 100; i++ {
		tree.Remove(ComparableInt(i))
	}
	assert.Equal(t, 0, tree.Size())
	assert.True(t, tree.IsEmpty())
}

func TestBinaryTreeContains(t *testing.T) {
	tree := binary_tree.New()

	tree.AddAll(
		ComparableInt(-1),
		ComparableInt(2),
		ComparableInt(3),
		ComparableInt(4),
		ComparableInt(5),
	)

	assert.True(t, tree.Contains(ComparableInt(-1)))
	assert.True(t, tree.Contains(ComparableInt(5)))
	assert.False(t, tree.Contains(ComparableInt(100)))
	assert.False(t, tree.Contains(ComparableInt(-500)))
}
