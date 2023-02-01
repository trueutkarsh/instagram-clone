package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"instagram-clone/internal/users"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	// Users
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK!")
	})
	r.POST("/users", users.HandleCreateUser(db))

}
