package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Dahs81/coffee_micro_service/models"
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
	// query object
	query := bson.M{}

	search := r.URL.Query()

	// TODO: Check what this should be queried by
	for _, name := range []string{"name", "price", "size"} {
		if search[name] != nil {
			query[name] = search[name][0]
		}
	}

	coffees := []models.Coffee{}

	// Find all by the query that is entered
	cc.Session.DB("coffee").C("coffees").Find(query).Iter().All(&coffees)

	// If only one is found, the return the object, else return an array of objects
	switch len(coffees) {
	case 0:
		fmt.Fprintf(w, "%s", errors.New("This is an error"))
	case 1:
		jsn, _ := json.Marshal(coffees[0])
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", jsn)
	default:
		jsnArr, _ := json.Marshal(coffees)
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", jsnArr)
	}
}

// CreateCoffee - Function that creates a new coffee drink
func (cc CoffeeController) CreateCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	c := models.Coffee{}

	json.NewDecoder(r.Body).Decode(&c)

	c.ID = bson.NewObjectId()

	cc.Session.DB("coffee").C("coffees").Insert(c)

	// Marshal converts Go struct to JSON
	jsn, _ := json.Marshal(c)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", jsn)
}

// GetCoffee - Function that retrieves a single coffee drink
func (cc CoffeeController) GetCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Lines 42 - 51 could be abstracted out
	// Get the id from the httprouter.Params - This returns a string
	id := params.ByName("id")

	// Make sure id is of type ObjectId
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Not found")
		return
	}

	// Get id
	cid := bson.ObjectIdHex(id)

	c := models.Coffee{}

	if err := cc.Session.DB("coffee").C("coffees").FindId(cid).One(&c); err != nil {
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
	// Get the id from the params
	id := params.ByName("id")

	// Make sure id is of type ObjectId
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Error: Not found")
		return
	}

	// Get id
	cid := bson.ObjectIdHex(id)

	// Get the old data and store it
	c := models.Coffee{}
	if err := cc.Session.DB("coffee").C("coffees").FindId(cid).One(&c); err != nil {
		w.WriteHeader(404)
		return
	}

	// Map that contains the new payload
	update := map[string]interface{}{}

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	r.Body.Close()

	// Check map for update
	if update["name"] != nil {
		c.Name = update["name"].(string)
	}

	if update["price"] != nil {
		c.Price = update["price"].(string)
	}

	if update["size"] != nil {
		c.Size = update["size"].(string)
	}

	cc.Session.DB("coffee").C("coffees").UpdateId(c.ID, c)

	// Marshal converts Go struct to JSON
	jsn, _ := json.Marshal(c)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsn)
}

// DeleteCoffee -
func (cc CoffeeController) DeleteCoffee(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Get id
	cid := bson.ObjectIdHex(id)

	if err := cc.Session.DB("coffee").C("coffees").RemoveId(cid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}
