package main

import (
	"fmt"
	"instagram-clone/internal/media"
	"instagram-clone/internal/routes"
	"instagram-clone/internal/users"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {

	// set up db
	db, err := setupDB()
	if err != nil {
		fmt.Printf("Connecting to db failed with err %s", err)
		log.Fatal(err)
	}
	defer db.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})

	if err != nil {
		fmt.Printf("Unable to create session with AWS %s", err)
		log.Fatal(err)
	}

	r := gin.Default()
	// r.Group("/api")
	routes.RegisterRoutes(r, db, sess)

	log.Fatal(r.Run("localhost:3000"))
}

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", getDBargs())
	if err != nil {
		return nil, err
	}

	//Migrations
	if err := db.AutoMigrate(
		&users.User{},
		&users.Follower{},
		&media.Media{},
	).Error; err != nil {
		return nil, err
	}

	return db, nil
}

func getDBargs() string {

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"0.0.0.0",
		"5432",
		"postgres",
		"instadb_dev",
		"password",
	)

}
