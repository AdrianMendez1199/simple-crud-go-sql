package models

import "time"

type Student struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdateAt  time.Time
}
