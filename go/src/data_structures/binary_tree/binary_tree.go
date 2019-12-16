package binary_tree

import "github.com/99_problems/go/src/data_structures"

type Comparable interface {
	LessThan(value data_structures.Any) bool
	GreaterThan(value data_structures.Any) bool
}

type Tree interface {
	Add(any Comparable)
	AddAll(any ...Comparable)
	Remove(any Comparable) bool
	Contains(any Comparable) bool
	Size() int
	IsEmpty() bool
}

type Node struct {
	Left  *Node
	Right *Node
	Value Comparable
}

func newNode(value Comparable) *Node {
	return &Node{Value: value}
}

type BinaryTree struct {
	Head *Node
}

func New() Tree {
	return &BinaryTree{}
}

func (tree *BinaryTree) AddAll(values ...Comparable) {
	for _, value := range values {
		tree.Add(value)
	}
}

func (tree *BinaryTree) Add(value Comparable) {
	if tree.IsEmpty() {
		tree.Head = newNode(value)
		return
	}

	add(tree.Head, value)
}

func add(node *Node, value Comparable) {
	// ==
	if node.Value == value {
		return
	}

	// <
	if value.LessThan(node.Value) {
		if node.Left == nil {
			node.Left = newNode(value)
			return
		}

		add(node.Left, value)
		return
	}

	// >
	if value.GreaterThan(node.Value) {
		if node.Right == nil {
			node.Right = newNode(value)
			return
		}

		add(node.Right, value)
		return
	}
}

func addAll(from *Node, to *Node) {
	if to == nil {
		panic("To node can't be nil")
	}

	if from == nil {
		return
	}

	if from.Left != nil {
		addAll(from.Left, to)
	}

	if from.Right != nil {
		addAll(from.Right, to)
	}
}

func (tree BinaryTree) Size() int {
	return size(tree.Head)
}

func size(node *Node) int {
	if node == nil {
		return 0
	}

	if hasBothLeafs(node) {
		return size(node.Left) + size(node.Right)
	}

	if node.Left != nil {
		return 1 + size(node.Left)
	}

	if node.Right != nil {
		return 1 + size(node.Right)
	}

	return 1

}

func (tree BinaryTree) IsEmpty() bool {
	if tree.Head == nil {
		return true
	}
	return false
}

func (tree *BinaryTree) Remove(value Comparable) bool {
	parent := findParent(tree.Head, value)

	if parent.isEmpty() {
		return false
	}

	if parent.pointToHead {
		tree.Head = nil
		return true
	}

	var old *Node
	if parent.node.Left != nil && parent.node.Left.Value == value {
		old = parent.node.Left
		parent.node.Left = nil
	} else {
		old = parent.node.Right
		parent.node.Right = nil
	}
	addAll(old, tree.Head)
	return true
}

type referenceToParent struct {
	node        *Node
	pointToHead bool
}

func referenceToHeader(node *Node) referenceToParent {
	return referenceToParent{node: node, pointToHead: true}
}

func emptyReferenceToParent() referenceToParent {
	return referenceToParent{}
}

func (r referenceToParent) isEmpty() bool {
	var empty referenceToParent
	return r == empty
}

// return: node, is point to header
func findParent(node *Node, value Comparable) referenceToParent {
	if node == nil {
		return emptyReferenceToParent()
	}

	if node.Value == value {
		return referenceToHeader(node)
	}

	if hasBothLeafs(node) {
		foundLeft := findParent(node.Left, value)
		if foundLeft.node != nil {
			return foundLeft
		}
		foundRight := findParent(node.Right, value)
		if foundRight.node != nil {
			return foundRight
		}
		return emptyReferenceToParent()
	}

	if node.Left != nil {
		if node.Left.Value == value {
			return referenceToParent{node: node}
		}
		return findParent(node.Left, value)
	}

	if node.Right != nil {
		if node.Right.Value == value {
			return referenceToParent{node: node}
		}
		return findParent(node.Right, value)
	}

	return emptyReferenceToParent()
}

func hasBothLeafs(node *Node) bool {
	return node.Left != nil && node.Right != nil
}

func (tree BinaryTree) Contains(value Comparable) bool {
	if tree.IsEmpty() {
		return false
	}

	return contains(tree.Head, value)
}

func contains(node *Node, value Comparable) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	}

	if value.LessThan(node.Value) {
		return contains(node.Left, value)
	} else {
		return contains(node.Right, value)
	}
}
