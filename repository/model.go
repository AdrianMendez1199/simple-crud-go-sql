package repository

import "time"

//base model, all application models inherit from this model
type Model struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
