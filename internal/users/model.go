package users

import (
	"instagram-clone/internal/utils"
	"time"
)

type User struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Handle         string     `json:"handle" gorm:"unique"`
	Email          string     `json:"email"`
	DateOfBirth    utils.Date `json:"dateOfBirth" gorm:"not null"`
	EmailVerfified bool       `json:"emailVerified"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Follower struct {
	FromUserID   uint `gorm:"not null"`
	TargetUserID uint `gorm:"not null"`
}
