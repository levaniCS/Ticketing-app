package routes

import (
	"example.com/ticketing/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Middleware will run for these routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	// Registration on event endpoints
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/cancel", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
