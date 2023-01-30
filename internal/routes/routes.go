package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Users
	r.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "trueutkarsh")
	})

}
