package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

// this function call repository to create teacher
func (a *API) createTeacher(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.teacherRepo)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	err = a.teacherRepo.CreateTeacher(a.teacherRepo)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&Response{
		Status:  "OK",
		Message: "",
	})
}
