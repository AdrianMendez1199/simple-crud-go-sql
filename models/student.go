package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/AdrianMendez1199/simple-crud-go-sql/database"
)

type Student struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdateAt  time.Time
}

// this function create student into database
func CreateStudent(s Student) error {
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

func GetStudents() (error, []Student) {
	query := `SELECT *FROM studens`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err, nil
	}

	result, err := stmt.Query()

	if err != nil {
		return err, nil
	}

	// for result.Next() {
	// 	fmt.Println("TRS")
	// }

	fmt.Println(result)
	return nil, []Student{
		{
			Name:   "TEST",
			Age:    21,
			Active: true,
		},
	}
}
