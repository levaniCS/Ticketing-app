package routes

import (
	"net/http"
	"strconv"

	"example.com/ticketing/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for an event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
	}

	var event models.Event
	event.ID = eventId

	err = event.CalcelRegistration(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel user for an event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event registration is canceled"})
}
