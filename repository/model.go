package repository

import "time"

type Model struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
