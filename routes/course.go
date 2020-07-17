package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createCourse(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a.courseRepo)

	err = a.courseRepo.CreateCourse(a.courseRepo)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Response{
			Status:  "NOK",
			Message: "An ocurrend Error creating course",
		})

	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&Response{
			Status:  "OK",
			Message: "Course created",
		})
	}
}

func (a *API) getCourseByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	studentRepo, err := a.courseRepo.GetCouseByID(id)

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

func (a *API) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := a.courseRepo.GetAllCourses()

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(Response{
			Status:  "NOK",
			Message: "An Ocurred Error",
		})

	} else {
		json.NewEncoder(w).Encode(courses)
	}
}
