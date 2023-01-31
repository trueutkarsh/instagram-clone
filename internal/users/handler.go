package users

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"secondName"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Handle      string    `json:"handle"`
}

func HandleCreateUser(c *gin.Context) {

}
