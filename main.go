package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

// HANDLER FUNCTIONS
func greet(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "<h1>Welcome to Koala Online Store</h1>")
}

func main() {
	router := mux.NewRouter()
	database.PopulatePlates()

	router.HandleFunc("/", greet).Methods("GET")

	router.HandleFunc("/plates", routes.GetAllPlates).Methods("GET")

	router.HandleFunc("/plates", routes.CreatePlate).Methods("POST")

	// app.use(/plates/:id)
	router.HandleFunc("/plates/{id}", routes.GetPlateById).Methods("GET")

	router.HandleFunc("/plates/{id}", routes.UpdatePlateById).Methods("PUT")

	router.HandleFunc("/plates/{id}", routes.DeletePlateById).Methods("DELETE")

	http.ListenAndServe(":8000", router)
	fmt.Println("Server running at port 8000")
}

// GOLANG IS A STATICALLY-TYPED LANGUAGE!!!

// MODEL
// CRUD service ==> Create, Read, Update, and Delete
// Create a new instance of a model
// Read data from the collection of that model
// Update models in the collection
// Delete ..

// Multiplexer (aka. the request handler)
// It receives incoming network requests -> http, ws, tcp, bluetooth
// based on the request, it determines the file to serve or the function to run

// HTTP
// Headers contain metadata that suggest to the web server how to handle the request
// Headers -> Request Type -> GET, POST, UPDATE, PATH, PUT
//         -> Content Type
// 		   -> Multipart Formdata
// User Agent
// HTTP Version ->

// Calls a handler function
