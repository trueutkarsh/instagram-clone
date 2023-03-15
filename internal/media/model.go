package media

import (
	"instagram-clone/internal/users"
	"time"
)

type Media struct {
	ID     uint `json:"id"`
	UserID uint `json:"userID"`
	User   users.User
	// UserID is foreign key
	Caption   string    `json:"caption"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"uploadTime"`
}
