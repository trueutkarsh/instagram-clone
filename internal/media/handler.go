package media

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func HandleCreateMedia(db *gorm.DB, sess *session.Session) gin.HandlerFunc {
	s3_svc := s3.New(sess)
	db_svc := New(db)
	return func(c *gin.Context) {

		userID := c.PostForm("userID")
		caption := c.PostForm("caption")
		file, err := c.FormFile("file")
		if userID == "" || caption == "" || err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "failed to parse input from request",
			})
			return
		}

		user_id, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error":       err.Error(),
				"description": "failed to parse input from request",
			})
			return
		}

		// open contents of the file to upload
		data, err := file.Open()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error":       err.Error(),
				"description": "unable to open received file",
			})
			return
		}
		defer data.Close()

		// upload file to S3
		filekey := file.Filename

		_, err = s3_svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String("instagram-clone-images-bucket"),
			Key:    &filekey,
			Body:   data,
			// ACL:    aws.String("public-read"),
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error":       err.Error(),
				"description": "unable to upload media to s3",
			})
			return
		}

		input := Media{
			UserID:  uint(user_id),
			Caption: caption,
			URL: fmt.Sprintf(
				"https://%s.s3-%s.amazonaws.com/%s",
				"instagram-clone-images-bucket",
				"ap-southeast-1",
				filekey),
		}

		media, err := db_svc.CreateItem(&input)
		if err != nil {
			// delete object from s3
			s3_svc.DeleteObject(&s3.DeleteObjectInput{
				Bucket: aws.String("instagram-clone-images-bucket"),
				Key:    &filekey,
			})

			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error":       err.Error(),
				"description": "unable to save media",
			})

			return
		}

		c.JSON(http.StatusCreated, map[string]string{
			"postID": strconv.FormatUint(uint64(media.ID), 10),
		})

	}
}
