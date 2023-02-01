package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Handle      string    `json:"handle"`
}

func HandleCreateUser(db *gorm.DB) gin.HandlerFunc {
	service := New(db)
	return func(c *gin.Context) {
		input := CreateUserInput{}
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "failed to parse input from request",
			})
			return
		}

		createdUser, err := service.CreateItem(&input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error":       err.Error(),
				"description": "failed to create user",
			})
			return
		}

		c.JSON(http.StatusCreated, createdUser)
	}

}
