package repository

import "github.com/AdrianMendez1199/simple-crud-go-sql/database"

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

// This function return all user into db
func (s *Student) GetStudents() ([]Student, error) {
	// Slice studens
	var students []Student

	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Where("active <> ?", false).Select("id, name, age").Find(&students)

	if err != nil {
		return students, err.Error
	}
	return students, nil
}

func (s *Student) GetStudentById(id string) (Student, error) {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	st := Student{}

	err := db.Where("active <> ?", false).Select("id, name, age").First(&st, id)

	if err != nil {
		return st, err.Error
	}

	return st, nil
}
