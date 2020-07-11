package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router http.Handler
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

	// Routes
	r.HandleFunc("/students", a.CreateStudent).Methods(http.MethodGet)

	a.router = r
	return a
}
