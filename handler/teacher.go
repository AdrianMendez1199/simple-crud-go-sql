package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createTeacher(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&a.teacherRepo)

	if err != nil {
		return
	}

	err = a.teacherRepo.CreateTeacher(*a.teacherRepo)
	rs := &Response{"OK", "Teacher creates"}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rs = &Response{"NOK", "Error creating teacher"}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&rs)

}

func (a *API) getTeacherByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	teacher, err := a.teacherRepo.GetTeacherByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&Response{"NOK", "Teacher not found"})
		return
	}

	json.NewEncoder(w).Encode(&teacher)

}

func (a *API) getTeachers(w http.ResponseWriter, r *http.Request) {

	teachers, err := a.teacherRepo.GetTeachers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Response{"NOK", "Error creating teacher"})
		return
	}

	json.NewEncoder(w).Encode(&teachers)

}
