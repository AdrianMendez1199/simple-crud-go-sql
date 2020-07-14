package repository

import "github.com/AdrianMendez1199/simple-crud-go-sql/database"

type Course struct {
	// Base model
	Model

	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"imgUrl,omitempty"`
	TeacherId   uint   `json:"teacherId"`
}

func (c *Course) CreateCourse(co *Course) error {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Save(&co)

	if err != nil {
		return err.Error
	}

	return nil
}
