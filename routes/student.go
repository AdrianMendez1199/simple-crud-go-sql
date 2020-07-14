package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var res = &Response{}

// create user into db
func (a *API) CreateStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	fmt.Println(a.studentRepo)
	err := decoder.Decode(&a.studentRepo)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(&Response{
			Status:  "NOK",
			Message: "Error creating user",
		})

	} else {

		a.studentRepo.CreateStudent(a.studentRepo)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(&Response{
			Status: "OK", Message: "user created",
		})

	}

}

func (a *API) GetStudents(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := a.studentRepo.GetStudents()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(&Response{
			Status:  "NOK",
			Message: "an Ocurred Error try later",
		})

	}

	json.NewEncoder(w).Encode(studentRepo)

}

func (a *API) GetStudentById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	studentRepo, err := a.studentRepo.GetStudentById(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(Response{
			Status:  "NOK",
			Message: "student not found",
		})

	} else {
		json.NewEncoder(w).Encode(studentRepo)
	}

}
