package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdrianMendez1199/simple-crud-go-sql/models"
	"github.com/AdrianMendez1199/simple-crud-go-sql/repository"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (a *API) CreateStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	res := &Response{}

	var s models.Student

	err := decoder.Decode(&s)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res = &Response{Status: "NOK", Message: "Error creating user"}
		json.NewEncoder(w).Encode(res)
		// panic(err)
	} else {
		// Create student in DB
		err = repository.CreateStudent(s)

		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusCreated)
		res = &Response{Status: "OK", Message: "user created"}
		json.NewEncoder(w).Encode(res)
	}

}
