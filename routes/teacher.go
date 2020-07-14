package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// this function call repository to create teacher
func (a *API) createTeacher(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.teacherRepo)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	err = a.teacherRepo.CreateTeacher(*a.teacherRepo)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Response{
			Status:  "NOK",
			Message: "Error creating teacher",
		})

	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&Response{
			Status:  "OK",
			Message: "Teacher create",
		})

	}

}

func (a *API) getTeacherByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	teacher, err := a.teacherRepo.GetTeacherByID(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(&Response{
			Status:  "NOK",
			Message: "Teacher not found",
		})
	} else {
		json.NewEncoder(w).Encode(&teacher)
	}

}

func (a *API) getTeachers(w http.ResponseWriter, r *http.Request) {

	teachers, err := a.teacherRepo.GetTeachers()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(&Response{
			Status:  "NOK",
			Message: "Error creating teacher",
		})

	} else {
		json.NewEncoder(w).Encode(&teachers)
	}

}
