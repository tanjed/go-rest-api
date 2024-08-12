package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanjed/go-rest-api/models"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	server := gin.Default()

	server.GET("/events", func(ctx *gin.Context) {
		events := models.GetAllEvents()
		response := Response{
			Success: true,
			Message: "Request success",
			Data:    events,
		}
		ctx.JSON(http.StatusOK, response)
	})

	server.POST("/events", func(ctx *gin.Context) {
		event := models.Event{}
		err := ctx.ShouldBindJSON(&event)

		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusUnprocessableEntity, Response{
				Success: false,
				Message: "Invalid data",
				Data:    []any{},
			})
			return
		}

		event.ID = 1
		event.CreadtedBy = 1
		event.Save()

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Event created",
			Data:    event,
		})
	})

	server.Run(":8080")
}
