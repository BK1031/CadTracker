package controller

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config"
	"server/model"
	"server/service"
	"strings"
)

func GetAllUsers(c *gin.Context) {
	result := service.GetAllUsers()
	c.JSON(http.StatusOK, result)
}

func GetUserByID(c *gin.Context) {
	result := service.GetUserByID(c.Param("userID"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found with given id: " + c.Param("userID")})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func CreateUser(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Set the user id to ensure that the user can only modify their own account
	input.ID = c.Param("userID")
	if err := service.CreateUser(input); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, service.GetUserByID(input.ID))
}

func DiscordLinkAccount(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(strings.Split(m.Content, " ")) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Command usage: `"+config.DiscordPrefix+"link [CadTracker Account ID]`")
		return
	}
	// First check if given userID exists
	userID := strings.Split(m.Content, config.DiscordPrefix+"link ")[1]
	user := service.GetUserByID(userID)
	if user.ID == "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "No CadTracker account found with that ID!")
		return
	}
	user.DiscordID = m.Author.ID
	_ = service.CreateUser(user)
	_, _ = s.ChannelMessageSend(m.ChannelID, "Successfully linked <@"+m.Author.ID+"> to "+user.FirstName+" "+user.LastName+" on CadTracker!")
}
