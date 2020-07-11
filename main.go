package main

import (
	"fmt"
	"log"

	"github.com/AdrianMendez1199/simple-crud-go-sql/models"
)

func main() {
	// students, err := models.GetStudents()

	student, err := models.GetUserById(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(student)

	// for _, student := range students {
	// }
	// student := models.Student{
	// 	Name:   "Adrian",
	// 	Age:    21,
	// 	Active: true,
	// }

	// err := models.CreateStudent(student)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("user created")
}
