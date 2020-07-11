package models

import (
	"time"
)

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
