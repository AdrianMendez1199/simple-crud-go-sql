package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var res = &Response{}

func (a *API) createStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.studentRepo)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

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

func (a *API) getStudents(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := a.studentRepo.GetStudents()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(&Response{
			Status:  "NOK",
			Message: "an Ocurred Error try later",
		})

	}

	json.NewEncoder(w).Encode(studentRepo)

}

func (a *API) getStudentByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	studentRepo, err := a.studentRepo.GetStudentByID(id)

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
