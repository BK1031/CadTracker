package controller

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config"
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

func DiscordStartEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// First check if user has a CadTracker account
	user := service.GetUserByID(m.Author.ID)
	if user.ID == "" {
		s.ChannelMessageSend(m.ChannelID, "You do not have a connected CadTracker account! Please create one at https://cad.bk1031.dev, or link your Discord account to your CadTracker account using the `"+config.DiscordPrefix+"link` command.")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Starting event...")
}
