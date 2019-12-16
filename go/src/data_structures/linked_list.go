package data_structures

import (
	"github.com/99_problems/go/src/utils"
)

type Node struct {
	Value    utils.Any
	Previous *Node
	Next     *Node
}

func (node Node) hasNext() bool {
	if node.Next == nil {
		return false
	}
	return true
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) IsEmpty() bool {
	if list.size == 0 {
		return true
	}
	return false
}

func (list *LinkedList) Add(obj utils.Any) {
	if list.IsEmpty() {
		node := &Node{Value: obj}
		list.head = node
		list.tail = node
	} else {
		node := &Node{Value: obj, Previous: list.tail}
		list.tail.Next = node
		list.tail = node
	}

	list.size += 1
}

func isNodeInTheMiddle(list *LinkedList, node *Node) bool {
	if node != list.head && node != list.tail {
		return true
	}
	return false
}

func (list *LinkedList) Remove(idx int) utils.Any {
	node := list.get(idx)
	return list.remove(node)
}

func (list *LinkedList) Filter(checker func(utils.Any) bool) List {
	newList := NewLinkedList()
	iterator := list.Iterator()

	for iterator.HasNext() {
		current := iterator.Next()
		if checker(current) {
			newList.Add(current)
		}
	}

	return newList
}

func (list *LinkedList) Find(checker func(utils.Any) bool) utils.Any {
	iterator := list.Iterator()

	for iterator.HasNext() {
		current := iterator.Next()
		if checker(current) {
			return current
		}
	}

	return nil
}

func (list *LinkedList) RemoveObject(obj utils.Any) utils.Any {
	node := list.find(func(value utils.Any) bool {
		if value == obj {
			return true
		}
		return false
	})
	return list.remove(node)
}

func (list *LinkedList) remove(node *Node) utils.Any {
	if node == nil {
		return nil
	}

	originalSize := list.size
	list.size -= 1

	// Head and Tail the same
	if originalSize == 1 {
		list.head = nil
		list.tail = nil
		return node.Value
	}

	if isNodeInTheMiddle(list, node) {
		priv := node.Previous
		next := node.Next

		priv.Next = next
		next.Previous = priv

		return node.Value
	}

	if node == list.head {
		list.head = node.Next
		return node.Value
	}

	if node == list.tail {
		list.tail = node.Previous
		return node.Value
	}

	return node.Value
}

func (list *LinkedList) find(predicate func(utils.Any) bool) *Node {
	if list.IsEmpty() {
		return nil
	}

	current := &Node{Next: list.head}

	for current.hasNext() {
		current = current.Next
		if predicate(current.Value) {
			return current
		}
	}

	return nil
}

func (list *LinkedList) get(idx int) *Node {
	if idx < 0 || idx >= list.size {
		panic("Index out of bound exception")
	}

	current := list.head

	for i := 0; i < idx; i++ {
		current = current.Next
	}

	return current
}

func (list *LinkedList) Get(idx int) utils.Any {
	return list.get(idx).Value
}

func (list *LinkedList) Size() int {
	return list.size
}

type linkedListIterator struct {
	nextNode *Node
}

func (it linkedListIterator) HasNext() bool {
	if it.nextNode.hasNext() {
		return true
	}

	return false
}

func (it *linkedListIterator) Next() utils.Any {
	if !it.nextNode.hasNext() {
		panic("Iterator is out of bounds")
	}
	it.nextNode = it.nextNode.Next
	return it.nextNode.Value
}

func (list LinkedList) Iterator() Iterator {
	return &linkedListIterator{&Node{Next: list.head}}
}
