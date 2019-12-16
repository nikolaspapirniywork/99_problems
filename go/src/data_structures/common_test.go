package data_structures_test

import (
	"github.com/99_problems/go/src/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

type Example struct {
	Age  int
	Name string
}

func TestToBytes(t *testing.T) {
	data_structures.ToBytes("asdasd")
	data_structures.ToBytes(10)
	data_structures.ToBytes(10.5)
}

func TestHashString(t *testing.T) {
	e1 := data_structures.HashObject("key1")
	e2 := data_structures.HashObject("key1")
	e3 := data_structures.HashObject("key2")
	assert.Equal(t, e1, e2)
	assert.NotEqual(t, e1, e3)
}

func TestHash(t *testing.T) {
	e1 := data_structures.HashObject(Example{21, "ivan"})
	e2 := data_structures.HashObject(Example{21, "ivan"})
	e3 := data_structures.HashObject(Example{21, "ivan1"})
	e4 := data_structures.HashObject(Example{22, "ivan"})

	log.Printf(strconv.Itoa(e1))
	assert.Equal(t, e1, e2)
	assert.NotEqual(t, e1, e3)
	assert.NotEqual(t, e1, e4)
	assert.NotEqual(t, e1, e4)
}

