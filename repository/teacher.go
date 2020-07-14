package repository

import "github.com/AdrianMendez1199/simple-crud-go-sql/database"

type Teacher struct {
	//inheriting from the base model
	Model

	Name     string `json:"name,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Age      int    `json:"age,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Active   bool   `json:"active,omitempty"`
}

func (tc *Teacher) CreateTeacher(t Teacher) error {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Save(&t)

	if err != nil {
		return err.Error
	}
	return nil
}

func (tc *Teacher) GetTeacherById(id string) (Teacher, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	t := Teacher{}
	err := db.Where("active = ?", true).Select("id, name, last_name, bio").First(&t, id)

	if err != nil {
		return t, err.Error
	}
	return t, nil
}

func (tc *Teacher) GetTeachers() ([]Teacher, error) {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	var teacher []Teacher

	err := db.Where("active = ?", true).Select("id, name, last_name, bio").Find(&teacher)

	if err != nil {
		return teacher, err.Error
	}
	return teacher, nil
}
