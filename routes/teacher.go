package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createTeacher(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.teacherRepo)

	if err != nil {
		log.Fatal(err)
	}

	err = a.teacherRepo.CreateTeacher(*a.teacherRepo)

	rs := &Response{"OK", "Teacher create"}

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		rs = &Response{"NOK", "Error creating teacher"}

	} else {

		w.WriteHeader(http.StatusCreated)

	}

	json.NewEncoder(w).Encode(&rs)

}

func (a *API) getTeacherByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	teacher, err := a.teacherRepo.GetTeacherByID(id)

	if err != nil {

		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(&Response{"NOK", "Teacher not found"})
	} else {

		json.NewEncoder(w).Encode(&teacher)
	}

}

func (a *API) getTeachers(w http.ResponseWriter, r *http.Request) {

	teachers, err := a.teacherRepo.GetTeachers()

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Response{"NOK", "Error creating teacher"})

	} else {
		json.NewEncoder(w).Encode(&teachers)
	}

}
