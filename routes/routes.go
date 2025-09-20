package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/events",getEvents)
	server.POST("/events",createEvent)
	server.GET("/events/:id",getEvent) // : dynamic example events/1,events/2 etc
	server.PUT("/events/:id",updateEvent)
	server.DELETE("/events/:id",deleteEvent)


}