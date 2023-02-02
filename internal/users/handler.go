package users

import (
	"fmt"
	"net/http"
	"strconv"
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

func HandleGetUser(db *gorm.DB) gin.HandlerFunc {
	service := New(db)
	return func(c *gin.Context) {
		user_id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "invalid user id",
			})
			return
		}

		item, err := service.GetItem(uint(user_id))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error":       err.Error(),
				"description": fmt.Sprintf("Failed to retrieve user with user id %d", user_id),
			})
			return
		}

		if item == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
				"error":       "User not found",
				"description": fmt.Sprintf("User with id %d does not exist", user_id),
			})
			return
		}

		c.JSON(http.StatusOK, item)

	}

}

func HandleUpdateItem(db *gorm.DB) gin.HandlerFunc {
	service := New(db)
	return func(c *gin.Context) {
		user_id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "invalid user id",
			})
			return
		}

		new_user := User{}
		err = c.ShouldBindJSON(&new_user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "failed to parse input from request",
			})
			return
		}
		// set primary key to zero value
		new_user.ID = 0
		err = service.UpdateItem(uint(user_id), &new_user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": fmt.Sprintf("Failed to update user with ID %d", user_id),
			})
			return
		}

		c.Status(http.StatusOK)

	}

}
