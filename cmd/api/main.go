package main

import (
	"fmt"
	"instagram-clone/internal/routes"
	"log"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// set up db
	db, err := setupDB()
	if err != nil {
		// fmt.Printf("Connecting to db failed with err %s", err)
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	// r.Group("/api")
	routes.RegisterRoutes(r)

	r.Run()
}

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", getDBargs())
	if err != nil {
		return nil, err
	}

	//Migrations
	if err := db.AutoMigrate(
		&user.User{},
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
