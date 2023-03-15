package media

import (
	"github.com/jinzhu/gorm"
)

type Service interface {
	CreateItem(item *Media) (*Media, error)
}

type Client struct {
	DB *gorm.DB
	// S3 *s3.S3
}

func New(db *gorm.DB) Service {
	//return &Client{DB: db, S3: s3}
	return &Client{DB: db}
}

func (client *Client) CreateItem(item *Media) (*Media, error) {
	return item, client.DB.Create(item).Error
}
