package data_structures

import (
	"github.com/99_problems/go/src/utils"
	"log"
)

const initialCapacity = 10

type ArrayList struct {
	arr      []utils.Any
	size     int
	capacity int
}

func NewArrayList() *ArrayList {
	return NewArrayListWithCapacity(initialCapacity)
}

func NewArrayListWithCapacity(capacity int) *ArrayList {
	initial := make([]utils.Any, capacity)

	return &ArrayList{
		arr:      initial,
		size:     0,
		capacity: capacity,
	}
}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) resize() {
	newCapacity := list.capacity * 2
	newArray := make([]utils.Any, newCapacity)

	copy(newArray[0:list.Size()], list.arr[0:list.Size()])

	list.arr = newArray
	list.capacity = newCapacity
}

func (list *ArrayList) Add(obj utils.Any) {
	if list.Size() == list.capacity {
		list.resize()
	}

	list.arr[list.Size()] = obj
	list.size += 1
}

func (list *ArrayList) Remove(idx int) utils.Any {
	oldValue := list.Get(idx)

	for i := idx; i < list.Size()-1; i++ {
		list.arr[i] = list.arr[i+1]
	}

	list.arr[list.Size()-1] = nil
	list.size -= 1

	return oldValue
}

func (list *ArrayList) Find(obj utils.Any) int {
	for i, value := range list.arr {
		if value == obj {
			return i
		}
	}

	return -1
}

func (list *ArrayList) IsEmpty() bool {
	if list.Size() == 0 {
		return true
	}
	return false
}

func (list *ArrayList) Get(idx int) utils.Any {
	if idx < 0 || idx >= list.capacity {
		panic("Index out of bound exception")
	}

	return list.arr[idx]
}

func (list *ArrayList) PrintAll() {
	for i, value := range list.arr {
		log.Printf("[%v] %v", i, value)
	}
}

type arrayListIterator struct {
	index int
	list  ArrayList
}

func (it arrayListIterator) HasNext() bool {
	if it.isWithinRange() {
		return true
	}

	return false
}

func (it *arrayListIterator) Next() utils.Any {
	if !it.isWithinRange() {
		panic("Iterator is out of bounds")
	}

	it.index += 1
	return it.list.Get(it.index)
}

func (it arrayListIterator) isWithinRange() bool {
	return it.index < it.list.Size()-1
}

func (list ArrayList) Iterator() Iterator {
	return &arrayListIterator{-1, list}
}

func (list *ArrayList) Filter(checker func(utils.Any) bool) List {
	var count int
	for i := 0; i < list.Size(); i++ {
		if checker(list.Get(i)) {
			count += 1
		}
	}

	newList := NewArrayListWithCapacity(count)
	for i := 0; i < list.Size(); i++ {
		if checker(list.Get(i)) {
			newList.Add(list.Get(i))
		}
	}

	return newList
}
