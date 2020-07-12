package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func (a *API) getTeacherById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	teacher, err := a.teacherRepo.GetTeacherById(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&teacher)

}
