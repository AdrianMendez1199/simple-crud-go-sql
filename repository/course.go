package repository

import (
	"fmt"
	"strings"

	"github.com/AdrianMendez1199/simple-crud-go-sql/database"
)

// Course model extends from base Model
type Course struct {
	// Base model
	Model

	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Image       string  `json:"imgUrl,omitempty"`
	Active      bool    `json:"active,omitempty"`
	TeacherID   uint    `json:"teacherId,omitempty"`
	Teacher     Teacher `json:"teacher,omitempty" gorm:"foreignkey:TeacherID"`
}

// CreateCourse creates course in database
func (c *Course) CreateCourse(co *Course) error {
	db := database.GetInstance().GetConnection()
	defer db.Close()

	err := db.Save(&co)

	if err != nil {
		return err.Error
	}

	return nil
}

// GetCouseByID return a course, based on the id passed by paramtro
func (c *Course) GetCouseByID(id string) (Course, error) {
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

// GetAllCourses return all courses that are active
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

// SearchCourse return courses
func (c *Course) SearchCourse(search string) ([]Course, error) {
	courses := []Course{}

	db := database.GetInstance().GetConnection()
	defer db.Close()

	search = strings.ToLower(search)

	fmt.Println(search)

	err := db.Where("lower(title) LIKE ? AND active = ?", "%"+search+"%", true).
		Preload("Teacher", "active = ?", true).
		Find(&courses)

	if err != nil {
		return courses, err.Error
	}

	return courses, nil
}
