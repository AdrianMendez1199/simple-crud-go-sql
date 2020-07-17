package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createStudent(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&a.studentRepo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{"NOK", "Bad request"})
		return
	}

	_, err = a.studentRepo.CreateStudent(a.studentRepo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Response{"NOK", "error creating user"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Response{"OK", "user created"})

}

func (a *API) getStudents(w http.ResponseWriter, r *http.Request) {

	students, err := a.studentRepo.GetStudents()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Response{"NOK", "an Ocurred Error try later"})
		return
	}

	json.NewEncoder(w).Encode(students)

}

func (a *API) getStudentByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	student, err := a.studentRepo.GetStudentByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&Response{"NOK", "student not found"})
		return
	}
	json.NewEncoder(w).Encode(student)

}
