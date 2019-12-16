package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

type Any = interface{}

type wrapper struct {
	Obj Any
}

func ToBytes(obj Any) []byte {
	if reflect.ValueOf(obj).Kind() != reflect.Struct {
		obj = wrapper{obj}
	}

	data, err := bson.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return data
}

func HashObject(obj Any) int {
	result := binary.BigEndian.Uint64(hash(ToBytes(obj)))
	return int(result)
}

func hash(data []byte) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}
