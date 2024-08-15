package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	resposegenerator "github.com/tanjed/go-rest-api/helpers/respose-generator"
	"github.com/tanjed/go-rest-api/models"
)

func Get(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resposegenerator.New(false, "Something went worng", []interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, resposegenerator.New(true, "Request success", events))
}

func Show(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := models.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, resposegenerator.New(false, "Invalid data", []interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, resposegenerator.New(true, "Successfully fetched", event))
}

func Store(ctx *gin.Context) {
	event := models.Event{}
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusUnprocessableEntity, resposegenerator.New(false, "Invalid data", []interface{}{}))
		return
	}

	event.CreadtedBy = 1
	event.Save()

	ctx.JSON(http.StatusOK, resposegenerator.New(true, "Event created", event))
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := models.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, resposegenerator.New(false, "Invalid data", []interface{}{}))
		return
	}
	event := models.Event{}
	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusUnprocessableEntity, resposegenerator.New(false, "Invalid data", []interface{}{}))
		return
	}

	event.ID, _ = strconv.ParseInt(id, 10, 64)

	isUpdated, err := event.Update()

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, resposegenerator.New(false, "Unable to update", []interface{}{}))
		return
	}

	if !isUpdated {
		ctx.JSON(http.StatusInternalServerError, resposegenerator.New(false, "Noting to update", []interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, resposegenerator.New(true, "Updated successfully", []interface{}{}))

}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := models.GetById(id)
	fmt.Println(event.ID)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, resposegenerator.New(false, "Invalid data", []interface{}{}))
		return
	}

	isUpdated, err := event.Delete()

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, resposegenerator.New(false, "Unable to delete", []interface{}{}))
		return
	}

	if !isUpdated {
		ctx.JSON(http.StatusInternalServerError, resposegenerator.New(false, "Noting to delete", []interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, resposegenerator.New(true, "Deleted successfully", []interface{}{}))

}
