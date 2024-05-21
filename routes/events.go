package routes

import (
	"net/http"
	"strconv"

	"example.com/ticketing/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "events not found"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "event not found"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "event not found"})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "event deleted successfully",
	})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindBodyWithJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "event not found"})
		return
	}

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindBodyWithJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "event updated successfully",
	})
}
