package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionAssets holds the name of the assets collection
	CollectionAssets = "assets"
)

// Assets model
type Asset struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Data string        `json:"data" bson:"data"`
}
