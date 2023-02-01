package users

import (
	"time"

	"github.com/jinzhu/gorm"
)

// services suppoprted for user

type Service interface {
	CreateItem(input *CreateUserInput) (*User, error)
}

type Client struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Service {
	return &Client{DB: db}
}

func (client *Client) CreateItem(input *CreateUserInput) (*User, error) {
	user := User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Handle:      input.Handle,
		DateOfBirth: input.DateOfBirth,
		Email:       input.Email,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return &user, client.DB.Create(&user).Error
}
