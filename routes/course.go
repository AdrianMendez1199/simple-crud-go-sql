package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *API) createCourse(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a.courseRepo)

	// Only for test
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

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
