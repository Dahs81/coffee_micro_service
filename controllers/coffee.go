package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dahs81/example_micro_service/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// CoffeeController - the controller for CRUD methods for some coffee
	CoffeeController struct {
		Session *mgo.Session
	}
)

// NewCoffeeController - returns a new controller of type CoffeeController
func NewCoffeeController(ses *mgo.Session) *CoffeeController {
	return &CoffeeController{ses}
}

// GetAllCoffee - Function that retrieves all the different coffee drinks
func (cc CoffeeController) GetAllCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// CreateCoffee - Function that creates a new coffee drink
func (cc CoffeeController) CreateCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	c := models.Coffee{}

	fmt.Println(c)
}

// GetCoffee - Function that retrieves a single coffee drink
func (cc CoffeeController) GetCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get the id from the httprouter.Params
	id := params.ByName("id")

	// Make sure id is of type ObjectId
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Get id
	convertedID := bson.ObjectIdHex(id)

	c := models.Coffee{}

	if err := cc.Session.DB("coffee").C("coffees").FindId(convertedID).One(&c); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal converts Go struct to JSON
	jsn, _ := json.Marshal(c)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsn)
}

// UpdateCoffee -
func (cc CoffeeController) UpdateCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// DeleteCoffee -
func (cc CoffeeController) DeleteCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
