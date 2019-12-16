package list

import (
	"github.com/99_problems/go/src/data_structures"
)

type List interface {
	Add(obj data_structures.Any)
	Remove(idx int) data_structures.Any
	Get(idx int) data_structures.Any
	IsEmpty() bool
	Size() int
	Iterator() Iterator
	Filter(checker func(data_structures.Any) bool) List
}

type Iterator interface {
	HasNext() bool
	Next() data_structures.Any
}
