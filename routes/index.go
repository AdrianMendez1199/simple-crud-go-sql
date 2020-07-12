package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router http.Handler
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Server interface {
	Router() http.Handler
}

func (a *API) Router() http.Handler {
	return a.router
}

func New() Server {
	a := &API{}

	r := mux.NewRouter()

	// Routes Studens
	r.HandleFunc("/students", a.CreateStudent).Methods(http.MethodPost)
	r.HandleFunc("/students", a.GetStudents).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", a.GetStudentById).Methods(http.MethodGet)

	a.router = r
	return a
}
