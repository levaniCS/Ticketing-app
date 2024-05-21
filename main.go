package main

import (
	"example.com/ticketing/db"
	"example.com/ticketing/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8000") // "localhost:8000"
}
