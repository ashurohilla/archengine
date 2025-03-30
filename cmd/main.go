package main

import (
	"archengine/internal/db"
	"archengine/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDatabase()
	r := gin.Default()
	routes.RegisterUserRoutes(r)
	r.Run(":8080")

}
