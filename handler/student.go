package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createStudent(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&a.studentRepo)

	if err != nil {
		response := newResponse(Error, "bad structure student", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	_, err = a.studentRepo.CreateStudent(a.studentRepo)

	if err != nil {
		response := newResponse(Error, "An ocurred error creating students", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "student created", nil)
	responseJSON(w, http.StatusCreated, response)

}

func (a *API) getStudents(w http.ResponseWriter, r *http.Request) {

	students, err := a.studentRepo.GetStudents()

	if err != nil {
		response := newResponse(Error, "Internal Server Error", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Error, "OK", students)
	responseJSON(w, http.StatusOK, response)

}

func (a *API) getStudentByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	student, err := a.studentRepo.GetStudentByID(id)

	if err != nil {
		response := newResponse(Error, "student not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse(Error, "OK", student)
	responseJSON(w, http.StatusOK, response)

}

func (a *API) updateStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&a.studentRepo); err != nil {
		response := newResponse(Error, "Invalid structure", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	student, err := a.studentRepo.UpdateStudent(id, *a.studentRepo)

	if err != nil {
		response := newResponse(Error, "student not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse(Message, "OK", student)
	responseJSON(w, http.StatusOK, response)

}
