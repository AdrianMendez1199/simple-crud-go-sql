package repository

import (
	"errors"

	"github.com/AdrianMendez1199/simple-crud-go-sql/database"
	"github.com/AdrianMendez1199/simple-crud-go-sql/models"
	"github.com/lib/pq"
)

// this function create student into database
func CreateStudent(s models.Student) error {
	query := `INSERT INTO studens (name, age, active)
						VALUES($1, $2, $3)`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(s.Name, s.Age, s.Active)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()

	if rows != 1 {
		return errors.New("Ocurred Errror")
	}

	return nil

}

// This function return all user into db
func GetStudents() ([]models.Student, error) {
	query := `SELECT id, name, age, active
						FROM studens`

	// Slice studens
	var students []models.Student

	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		s := models.Student{}

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Age,
			&s.Active,
		)

		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

func GetUserById(id int) (models.Student, error) {
	query := `SELECT name, age, active, create_at, update_at
						FROM studens WHERE id = $1`

	student := models.Student{}

	timeNull := pq.NullTime{}

	db := database.GetConnection()
	defer db.Close()

	row, err := db.Query(query, id)
	defer row.Close()

	if err != nil {
		return student, err
	}

	for row.Next() {

		err := row.Scan(
			&student.Name,
			&student.Age,
			&student.Active,
			&student.CreatedAt,
			&timeNull,
		)

		student.UpdatedAt = timeNull.Time

		if err != nil {
			return student, err
		}
	}

	return student, nil
}
