package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tanjed/go-rest-api/controllers"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", controllers.Get)
	server.POST("/events", controllers.Store)
	server.GET("/events/:id", controllers.Show)
	server.PUT("/events/:id", controllers.Update)
	server.DELETE("/events/:id", controllers.Delete)
}
