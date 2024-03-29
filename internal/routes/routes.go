package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"instagram-clone/internal/media"
	"instagram-clone/internal/users"

	"github.com/aws/aws-sdk-go/aws/session"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, sess *session.Session, aws_env map[string]string) {

	// Users
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK!")
	})
	r.POST("/users", users.HandleCreateUser(db))

	r.GET("/users/:user_id", users.HandleGetUser(db))
	r.PATCH("/users/:user_id", users.HandleUpdateItem(db))
	r.PUT("/users/:user_id", users.HandleUpdateItem(db))

	r.POST("/users/:user_id/follow", users.HandleFollowUser(db))
	r.POST("/users/:user_id/unfollow", users.HandleUnfollowUser(db))

	r.POST("/media", media.HandleCreateMedia(db, sess, aws_env))

}
