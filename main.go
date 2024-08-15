package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanjed/go-rest-api/database"
	"github.com/tanjed/go-rest-api/routes"
)

func main() {
	database.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
