package repository

import (
	"errors"
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
	query := `INSERT INTO teacher(name, lastname, age, bio)
						VALUES($1, $2, $3, $4)`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(t.Name, t.LastName, t.Age, t.Bio)

	if err != nil {
		return err
	}

	row, _ := result.RowsAffected()

	if row != 1 {
		return errors.New("Error creating teacher")
	}

	return nil
}
