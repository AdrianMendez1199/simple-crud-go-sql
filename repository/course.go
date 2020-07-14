package repository

import "github.com/AdrianMendez1199/simple-crud-go-sql/database"

type Course struct {
	// Base model
	Model

	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Image       string  `json:"imgUrl,omitempty"`
	Active      bool    `json:"active,omitempty"`
	TeacherId   uint    `json:"teacherId,omitempty"`
	Teacher     Teacher `json:"teacher,omitempty" gorm:"foreignkey:TeacherId"`
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

func (c *Course) GetCouseById(id string) (Course, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	co := Course{}
	err := db.Where("active = ?", true).Select("title, description, image, teacher_id").First(&co, id)
	//Load relation with teacher
	db.Model(co).Select("name, last_name, bio").Related(&co.Teacher)

	if err != nil {
		return co, err.Error
	}

	return co, nil
}

func (c *Course) GetAllCourses() ([]Course, error) {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	courses := []Course{}

	err := db.Preload("Teacher", "active = ?", true).
		Find(&courses)

	if err != nil {
		return courses, err.Error
	}

	return courses, nil
}
