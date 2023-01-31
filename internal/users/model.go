package users

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Handle      string    `json:"handle" gorm:"unique"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
