package routes

import (
	"encoding/json"
	"go-rest-api/database"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePlate(res http.ResponseWriter, req *http.Request) {

}

func GetAllPlates(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(database.Plates)
}

func GetPlateById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, plate := range database.Plates {
		if plate.Id == params["id"] {
			json.NewEncoder(res).Encode(plate)
			break
		}
	}
	json.NewEncoder(res).Encode("No item found!")
}
func UpdatePlateById(res http.ResponseWriter, req *http.Request) {

}
func DeletePlateById(res http.ResponseWriter, req *http.Request) {

}
