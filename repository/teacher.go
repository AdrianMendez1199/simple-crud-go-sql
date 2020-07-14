package repository

import "github.com/AdrianMendez1199/simple-crud-go-sql/database"

type Teacher struct {
	//inheriting from the base model
	Model

	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Age      int    `json:"age"`
	Bio      string `json:"bio"`
}

func (tc *Teacher) CreateTeacher(t Teacher) error {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	db.Save(&t)
	return nil
}

func (tc *Teacher) GetTeacherById(id string) (t Teacher, err error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	db.Select("id, name, last_name, bio").First(&t, id)
	return t, nil
}

func (tc *Teacher) GetTeachers() (t []Teacher, err error) {

	db := database.GetInstance().GetConnection()
	defer db.Close()

	db.Select("id, name, last_name, bio").Find(&t)
	return t, nil
}
