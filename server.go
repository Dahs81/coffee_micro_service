package main

import (
	"net/http"

	"github.com/Dahs81/coffee_micro_service/controllers"
	"github.com/Dahs81/simple-mgo-db/db"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// What I want to do
	// 1.  Create DB
	// 2.  Create router
	// 3.  Create controller and routes (possibly abstract this)
	// 4.  Setup http server and listen

	// TODO: Add negroni for middleware

	// Setup and start the mongoDB connection
	d := db.New()
	// d.SetEnv(db.Env{Host: "MONGO_HOST_TEST", Port: "MONGO_PORT_TEST"})\
	d.SetHost("127.0.0.1")
	d.SetPort(27017)
	d.Start("coffee") // DB name

	router := httprouter.New()

	coffcntlr := controllers.NewCoffeeController(d.Session)

	// Abstract this - Create a folder called routes with file called coffee.go
	// Maybe pass the routes to the controller or vise versa - need to think about this
	router.GET("/coffee/:id", coffcntlr.GetCoffee)
	router.GET("/coffee", coffcntlr.GetAllCoffee)
	router.POST("/coffee", coffcntlr.CreateCoffee)
	router.PUT("/coffee/:id", coffcntlr.UpdateCoffee)
	router.DELETE("/coffee/:id", coffcntlr.DeleteCoffee)

	http.ListenAndServe("localhost:3000", router)
}
