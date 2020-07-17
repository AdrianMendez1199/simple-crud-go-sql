package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) createStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.studentRepo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(&Response{"NOK", "Bad request"})

	} else {

		_, err := a.studentRepo.CreateStudent(a.studentRepo)

		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Response{"NOK", "error creating user"})

		} else {

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&Response{"OK", "user created"})

		}
	}

}

func (a *API) getStudents(w http.ResponseWriter, r *http.Request) {

	students, err := a.studentRepo.GetStudents()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Response{"NOK", "an Ocurred Error try later"})
		log.Println("error", err)
	} else {
		json.NewEncoder(w).Encode(students)
	}

}

func (a *API) getStudentByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	student, err := a.studentRepo.GetStudentByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&Response{"NOK", "student not found"})

		log.Println(err)
	} else {

		json.NewEncoder(w).Encode(student)

	}

}
