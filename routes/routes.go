package routes

import (
	"net/http"

	"github.com/AdrianMendez1199/simple-crud-go-sql/repository"
	"github.com/gorilla/mux"
)

type API struct {
	router   http.Handler
	userRepo *repository.Student
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

func initServices() *API {
	//User Repo
	u := new(repository.Student)

	return &API{userRepo: u}
}

func New() Server {

	a := &API{}
	initServices()

	r := mux.NewRouter()

	// Routes Studens
	r.HandleFunc("/students", a.CreateStudent).Methods(http.MethodPost)
	r.HandleFunc("/students", a.GetStudents).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", a.GetStudentById).Methods(http.MethodGet)

	a.router = r
	return a
}
