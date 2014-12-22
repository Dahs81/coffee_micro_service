package models

import "gopkg.in/mgo.v2/bson"

type (
	// Coffee - An object that stores data for a coffee drink
	Coffee struct {
		ID    bson.ObjectId `json:"id" bson:"_id"`
		Name  string        `json:"name" bson:"name"`
		Size  string        `json:"size" bson:"size"`
		Price string        `json:"price" bson:"price"`
	}
)

// Define any behavior here if there were anything that a Coffee would need to do.
