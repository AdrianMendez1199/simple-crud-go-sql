package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var res = &Response{}

// create user into db
func (a *API) CreateStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&a.studentRepo)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res = &Response{Status: "NOK", Message: "Error creating user"}
		json.NewEncoder(w).Encode(res)
		// panic(err)
	} else {
		// Create student in DB
		// wg := sync.WaitGroup{}
		// wg.Add(1)
		for i := 0; i < 10000000; i++ {
			// defer wg.Done()
			go a.studentRepo.CreateStudent(a.studentRepo)
		}

		// wg.Wait()

		w.WriteHeader(http.StatusCreated)
		res = &Response{Status: "OK", Message: "user created"}
		json.NewEncoder(w).Encode(res)
	}

}

func (a *API) GetStudents(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := a.studentRepo.GetStudents()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res = &Response{Status: "NOK", Message: "error testing"}
		json.NewEncoder(w).Encode(studentRepo)
	} else {
		json.NewEncoder(w).Encode(studentRepo)
	}
}

func (a *API) GetStudentById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	studentRepo, err := a.studentRepo.GetUserById(id)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res = &Response{Status: "NOK", Message: "error testing"}
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(studentRepo)
	}

}
