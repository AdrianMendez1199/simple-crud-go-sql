package repository

import (
	"github.com/AdrianMendez1199/simple-crud-go-sql/database"
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	ID     int    `json:"id"`
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

func (s *Student) GetUserById(id string) (Student, error) {
	// query := `SELECT name, age, active, create_at, update_at
	// 					FROM studens WHERE id = $1`

	student := Student{}

	// timeNull := pq.NullTime{}

	// db := database.GetInstance().GetConnection()
	// defer db.Close()

	// row, err := db.Query(query, id)
	// defer row.Close()

	// if err != nil {
	// 	return student, err
	// }

	// for row.Next() {

	// 	err := row.Scan(
	// 		&student.Name,
	// 		&student.Age,
	// 		&student.Active,
	// 		&student.CreatedAt,
	// 		&timeNull,
	// 	)

	// 	student.UpdatedAt = timeNull.Time

	// 	if err != nil {
	// 		return student, err
	// 	}
	// }

	return student, nil
}
