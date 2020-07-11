package main

import (
	"fmt"

	"github.com/AdrianMendez1199/simple-crud-go-sql/models"
)

func main() {
	fmt.Println(models.GetStudents())
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
