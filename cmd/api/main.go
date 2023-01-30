package main

import (
	"instagram-clone/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// r.Group("/api")
	routes.RegisterRoutes(r)

	r.Run()
}
