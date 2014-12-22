// Package models - Stores the model data for all model types and their behavior
// Currently there is just a single model type that has no behavior, but more could be added if needed
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
