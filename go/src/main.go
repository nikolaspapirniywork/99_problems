package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {
	data, err := bson.Marshal(&Person{Name:"Bob"})
}
