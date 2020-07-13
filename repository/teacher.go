package repository

import (
	"time"

	"github.com/AdrianMendez1199/simple-crud-go-sql/database"
)

type Teacher struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Age       int       `json:"age"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (tc *Teacher) CreateTeacher(t *Teacher) error {
	// query := `INSERT INTO teacher(name, lastname, age, bio)
	// 					VALUES($1, $2, $3, $4)`

	db := database.GetInstance().GetConnection()
	defer db.Close()

	// stmt, err := db.Prepare(query)

	// if err != nil {
	// 	return err
	// }

	// defer stmt.Close()

	// result, err := stmt.Exec(t.Name, t.LastName, t.Age, t.Bio)

	// if err != nil {
	// 	return err
	// }

	// row, _ := result.RowsAffected()

	// if row != 1 {
	// 	return errors.New("Error creating teacher")
	// }

	return nil
}

func (tc *Teacher) GetTeacherById(id string) (*Teacher, error) {
	// query := `SELECT name, lastname, age
	// 				 FROM teacher
	// 				 WHERE id = $1`

	// db := database.GetInstance().GetConnection()
	// defer db.Close()

	t := Teacher{}
	// rows, err := db.Query(query, id)
	// defer rows.Close()

	// if err != nil {
	// 	return nil, err
	// }

	// count := 0
	// for rows.Next() {
	// 	count++

	// 	err := rows.Scan(
	// 		&t.Name,
	// 		&t.LastName,
	// 		&t.Age,
	// 	)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &t, nil
}

func (tc *Teacher) GetTeachers() ([]Teacher, error) {
	// query := `SELECT name, lastname, age
	// 				FROM teacher`

	var teachers []Teacher

	db := database.GetInstance().GetConnection()
	defer db.Close()

	// rows, err := db.Query(query)
	// defer rows.Close()

	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	tch := Teacher{}

	// 	err := rows.Scan(
	// 		&tch.Name,
	// 		&tch.LastName,
	// 		&tch.Age,
	// 	)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	teachers = append(teachers, tch)
	// }

	return teachers, nil
}
