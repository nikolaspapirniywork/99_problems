package data_structures

import "github.com/99_problems/go/src/utils"

type List interface {
	Add(obj utils.Any)
	Remove(idx int) utils.Any
	Get(idx int) utils.Any
	IsEmpty() bool
	Size() int
	Iterator() Iterator
	Filter(checker func(utils.Any) bool) List
}

type Iterator interface {
	HasNext() bool
	Next() utils.Any
}
