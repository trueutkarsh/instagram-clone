package users

import (
	"time"

	"github.com/jinzhu/gorm"
)

// services suppoprted for user

type Service interface {
	CreateItem(input *CreateUserInput) (*User, error)
	GetItem(user_id uint) (*User, error)
	UpdateItem(user_id uint, new_user *User) error
	FollowUser(from uint, target uint) error
	UnfollowUser(from uint, target uint) error
}

type Client struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Service {
	return &Client{DB: db}
}

func (client *Client) CreateItem(input *CreateUserInput) (*User, error) {
	user := User{
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Handle:         input.Handle,
		DateOfBirth:    input.DateOfBirth,
		Email:          input.Email,
		EmailVerfified: false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return &user, client.DB.Create(&user).Error
}

func (client *Client) GetItem(user_id uint) (*User, error) {
	user := User{}
	return &user, client.DB.Find(&user, user_id).Error
}

func (client *Client) UpdateItem(user_id uint, new_user *User) error {
	old_user := User{ID: user_id}
	return client.DB.Model(&old_user).Updates(new_user).Error
}

func (client *Client) FollowUser(from uint, target uint) error {
	follower := Follower{
		FromUserID:   from,
		TargetUserID: target,
	}
	return client.DB.Create(follower).Error
}

func (client *Client) UnfollowUser(from uint, target uint) error {
	follower := Follower{
		FromUserID:   from,
		TargetUserID: target,
	}
	return client.DB.Delete(follower).Error
}
