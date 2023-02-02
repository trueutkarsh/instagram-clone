package main

import (
	"fmt"
	"instagram-clone/internal/routes"
	"instagram-clone/internal/users"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// set up db
	db, err := setupDB()
	if err != nil {
		fmt.Printf("Connecting to db failed with err %s", err)
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	// r.Group("/api")
	routes.RegisterRoutes(r, db)

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
