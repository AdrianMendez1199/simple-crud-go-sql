package repository

import (
	"github.com/AdrianMendez1199/simple-crud-go-sql/database"
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(100)"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

// this function create student into database
func (st *Student) CreateStudent(s *Student) (bool, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	db.Save(&s)
	return true, nil
}

// This function return all user into db
func (s *Student) GetStudents() ([]Student, error) {
	// Slice studens
	var students []Student

	db := database.GetInstance().GetConnection()
	defer db.Close()

	db.Select("name, age").Find(&students)

	return students, nil
}

func (s *Student) GetUserById(id string) (student Student, err error) {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	db.Select("name, age").First(&student, id)
	return student, nil
}
