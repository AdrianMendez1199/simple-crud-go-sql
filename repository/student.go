package repository

import (
	"errors"
	"time"

	"github.com/AdrianMendez1199/simple-crud-go-sql/database"
	"github.com/lib/pq"
)

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// this function create student into database
func (st *Student) CreateStudent(s *Student) error {
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
func (s *Student) GetStudents() ([]Student, error) {
	query := `SELECT id, name, age, active
						FROM studens`

	// Slice studens
	var students []Student

	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		st := Student{}

		err := rows.Scan(
			&st.ID,
			&st.Name,
			&st.Age,
			&st.Active,
		)

		if err != nil {
			return nil, err
		}
		students = append(students, st)
	}

	return students, nil
}

func (s *Student) GetUserById(id string) (Student, error) {
	query := `SELECT name, age, active, create_at, update_at
						FROM studens WHERE id = $1`

	student := Student{}

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
