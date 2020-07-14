package repository

import "github.com/AdrianMendez1199/simple-crud-go-sql/database"

// Represent a Student and inheriting from base Model
type Student struct {
	//inheriting from the base model
	Model

	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Active bool   `json:"active,omitempty"`
}

// this function create student into database
func (st *Student) CreateStudent(s *Student) (bool, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Save(&s)

	if err != nil {
		return false, err.Error
	}

	return true, nil
}

// This function return all student into db with state A
func (st *Student) GetStudents() ([]Student, error) {
	// Slice studens
	students := []Student{}

	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Where("active <> ?", false).Select("id, name, age").Find(&students)

	if err != nil {
		return students, err.Error
	}
	return students, nil
}

// This function return student based on an id
func (st *Student) GetStudentByID(id string) (Student, error) {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	s := Student{}

	err := db.Where("active <> ?", false).Select("id, name, age").First(&s, id)

	if err != nil {
		return s, err.Error
	}

	return s, nil
}
