package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionAssets holds the name of the assets collection
	CollectionExamples = "examples"
)

// Assets model
type Asset struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Data string        `json:"data" bson:"data"`
}
type Tx struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Data string        `json:"data" bson:"data"`
}

