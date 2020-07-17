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

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{"NOK", "Bad Request"})

	} else {
		err = a.courseRepo.CreateCourse(a.courseRepo)

		if err != nil {
			json.NewEncoder(w).Encode(&Response{"NOK", "Error creating course"})
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&Response{"OK", "Course created"})
		}
	}
}

func (a *API) getCourseByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	course, err := a.courseRepo.GetCouseByID(id)

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&Response{"NOK", "student not found"})

	} else {
		json.NewEncoder(w).Encode(course)
	}

}

func (a *API) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := a.courseRepo.GetAllCourses()

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{"NOK", "An Ocurred Error"})

	} else {
		json.NewEncoder(w).Encode(courses)
	}
}
