package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/service"
)

func GetAllEvents(c *gin.Context) {
	result := service.GetAllEvents()
	c.JSON(http.StatusOK, result)
}

func GetEventByID(c *gin.Context) {
	result := service.GetEventByID(c.Param("eventID"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No event found with given id: " + c.Param("eventID")})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func CreateEvent(c *gin.Context) {
	var input model.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Set the user id to ensure that the user can only modify their own account
	input.ID = c.Param("eventID")
	if err := service.CreateEvent(input); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, service.GetEventByID(input.ID))
}
