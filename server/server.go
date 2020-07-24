package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AdrianMendez1199/simple-crud-go-sql/handler"
)

// Start run http server
func Start() {

	fmt.Println("Server start")

	PORT := os.Getenv("API_PORT")

	if PORT == "" {
		PORT = "3000"
	}

	router := handler.New()

	s := &http.Server{
		Addr:         ":" + PORT,
		Handler:      router.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
