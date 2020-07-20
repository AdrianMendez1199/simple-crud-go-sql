package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createCourse(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&a.courseRepo)

	if err != nil {
		response := newResponse(Error, "Bad Request", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = a.courseRepo.CreateCourse(a.courseRepo)

	if err != nil {
		response := newResponse(Error, "error creating course", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Course creates", nil)
	responseJSON(w, http.StatusCreated, response)

}

func (a *API) getCourseByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	course, err := a.courseRepo.GetCouseByID(id)

	if err != nil {
		response := newResponse(Error, "error creating course", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse(Error, "OK", course)
	responseJSON(w, http.StatusOK, response)

}

func (a *API) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := a.courseRepo.GetAllCourses()

	if err != nil {
		response := newResponse(Error, "Internal Server Error", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", courses)
	responseJSON(w, http.StatusOK, response)
}

func (a *API) searchCourse(w http.ResponseWriter, r *http.Request) {

	params, ok := r.URL.Query()["search"]

	if !ok {
		response := newResponse(Error, "Search param is not provider", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	courses, err := a.courseRepo.SearchCourse(params[0])

	if err != nil {
		response := newResponse(Error, "course not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse(Message, "OK", courses)
	responseJSON(w, http.StatusOK, response)

}
