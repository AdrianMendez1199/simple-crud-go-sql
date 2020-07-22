package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createTeacher(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&a.teacherRepo)

	if err != nil {
		response := newResponse(Error, "Error in struct teacher", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = a.teacherRepo.CreateTeacher(*a.teacherRepo)

	if err != nil {
		response := newResponse(Error, "Error creating teacher", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", "Teacher creates")
	responseJSON(w, http.StatusOK, response)

}

func (a *API) getTeacherByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	teacher, err := a.teacherRepo.GetTeacherByID(id)

	if err != nil {
		response := newResponse(Error, "Teacher not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse(Message, "OK", &teacher)
	responseJSON(w, http.StatusOK, response)
}

func (a *API) getTeachers(w http.ResponseWriter, r *http.Request) {

	teachers, err := a.teacherRepo.GetTeachers()

	if err != nil {
		response := newResponse(Error, "Internal Server Error", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", &teachers)
	responseJSON(w, http.StatusOK, response)
}

func (a *API) updateTeacher(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&a.teacherRepo)

	if err != nil {
		response := newResponse(Error, "Invalid structure", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	teacher, err := a.teacherRepo.UpdateTeacher(id, *a.teacherRepo)

	if err != nil {
		response := newResponse(Error, "Internal Server Error", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", &teacher)
	responseJSON(w, http.StatusOK, response)

}
