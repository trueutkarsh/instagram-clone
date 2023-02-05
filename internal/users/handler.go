package users

import (
	"fmt"
	"net/http"
	"strconv"

	"instagram-clone/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
	FirstName   string     `json:"firstName" binding:"required" `
	LastName    string     `json:"lastName" binding:"required" `
	Email       string     `json:"email" binding:"required" `
	DateOfBirth utils.Date `json:"dateOfBirth" binding:"required" `
	Handle      string     `json:"handle" binding:"required"`
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

type FollowActionInput struct {
	TargetUserID string `json:"targetUserID"`
}

func HandleFollowUser(db *gorm.DB) gin.HandlerFunc {
	service := New(db)
	return func(c *gin.Context) {
		from_user_id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "invalid user id",
			})
			return
		}

		input := FollowActionInput{}
		err = c.ShouldBindJSON(&input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "failed to parse input from request",
			})
			return
		}

		target_user_id, err := strconv.ParseUint(input.TargetUserID, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "invalid target user id",
			})
			return
		}

		err = service.FollowUser(uint(from_user_id), uint(target_user_id))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": fmt.Sprintf("user %d unable to follow user %d", from_user_id, target_user_id),
			})
			return
		}

		c.Status(http.StatusOK)

	}
}

func HandleUnfollowUser(db *gorm.DB) gin.HandlerFunc {
	service := New(db)
	return func(c *gin.Context) {
		from_user_id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "invalid user id",
			})
			return
		}

		input := FollowActionInput{}
		err = c.ShouldBindJSON(&input)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "failed to parse input from request",
			})
			return
		}

		target_user_id, err := strconv.ParseUint(input.TargetUserID, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "invalid target user id",
			})
			return
		}

		err = service.UnfollowUser(uint(from_user_id), uint(target_user_id))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": fmt.Sprintf("user %d unable to follow user %d", from_user_id, target_user_id),
			})
			return
		}

		c.Status(http.StatusOK)

	}
}
