package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdrianMendez1199/simple-crud-go-sql/models"
	"github.com/AdrianMendez1199/simple-crud-go-sql/repository"
	"github.com/gorilla/mux"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var res = &Response{}
var s models.Student

// create user into db
func (a *API) CreateStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&s)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res = &Response{Status: "NOK", Message: "Error creating user"}
		json.NewEncoder(w).Encode(res)
		// panic(err)
	} else {
		// Create student in DB
		err = repository.CreateStudent(s)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusCreated)
		res = &Response{Status: "OK", Message: "user created"}
		json.NewEncoder(w).Encode(res)
	}

}

func (a *API) GetStudents(w http.ResponseWriter, r *http.Request) {
	response, err := repository.GetStudents()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res = &Response{Status: "NOK", Message: "error testing"}
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(response)
	}
}

func (a *API) GetStudentById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	response, err := repository.GetUserById(id)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res = &Response{Status: "NOK", Message: "error testing"}
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(response)
	}

}
