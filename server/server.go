package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AdrianMendez1199/simple-crud-go-sql/routes"
)

func Start() {

	fmt.Println("Server start")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	router := routes.New()

	s := &http.Server{
		Addr:         ":" + PORT,
		Handler:      router.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
