package utils_test

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
	utils.ToBytes("asdasd")
	utils.ToBytes(10)
	utils.ToBytes(10.5)
}

func TestHashString(t *testing.T) {
	e1 := utils.HashObject("key1")
	e2 := utils.HashObject("key1")
	e3 := utils.HashObject("key2")
	assert.Equal(t, e1, e2)
	assert.NotEqual(t, e1, e3)
}

func TestHash(t *testing.T) {
	e1 := utils.HashObject(Example{21, "ivan"})
	e2 := utils.HashObject(Example{21, "ivan"})
	e3 := utils.HashObject(Example{21, "ivan1"})
	e4 := utils.HashObject(Example{22, "ivan"})

	log.Printf(strconv.Itoa(e1))
	assert.Equal(t, e1, e2)
	assert.NotEqual(t, e1, e3)
	assert.NotEqual(t, e1, e4)
	assert.NotEqual(t, e1, e4)
}
