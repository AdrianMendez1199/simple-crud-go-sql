package main

import (
	"fmt"
	"log"

	"github.com/AdrianMendez1199/simple-crud-go-sql/models"
)

func main() {

	student := models.Student{
		Name:   "Adrian",
		Age:    21,
		Active: true,
	}

	err := models.CreateStudent(student)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user created")
}
