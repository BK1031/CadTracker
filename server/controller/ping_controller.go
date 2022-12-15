package controller

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CadTracker Server v" + config.Version + " is online!"})
}

func DiscordPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "CadTracker Server v"+config.Version+" is online!")
}
