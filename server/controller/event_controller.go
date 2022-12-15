package controller

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config"
	"server/model"
	"server/service"
	"strconv"
	"time"
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
	// Check if user already has an event running
	lastEvent := service.GetLastEventForUser(user.ID)
	if lastEvent.ID != "" && lastEvent.Start == lastEvent.Stop {
		_, _ = s.ChannelMessageSend(m.ChannelID, "It looks like you already have a recording running, did you mean to use the `"+config.DiscordPrefix+"stop` command?")
		return
	}
	// Create new event
	event := model.Event{
		ID:         strconv.FormatInt(time.Now().Unix(), 10),
		UserID:     user.ID,
		Start:      time.Now(),
		Stop:       time.Now(),
		Notes:      "",
		Orgasm:     false,
		Ejaculated: false,
		UpdatedAt:  time.Time{},
		CreatedAt:  time.Time{},
	}
	err := service.CreateEvent(event)
	if err != nil {
		return
	}
	s.ChannelMessageSend(m.ChannelID, "Starting event...")
	// TODO: Add event to subscription queue
	return
}

func DiscordStopEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	// First check if user has a CadTracker account
	user := service.GetUserByID(m.Author.ID)
	if user.ID == "" {
		s.ChannelMessageSend(m.ChannelID, "You do not have a connected CadTracker account! Please create one at https://cad.bk1031.dev, or link your Discord account to your CadTracker account using the `"+config.DiscordPrefix+"link` command.")
		return
	}
	// Check if user already has an event running
	lastEvent := service.GetLastEventForUser(user.ID)
	if lastEvent.ID == "" || lastEvent.Start != lastEvent.Stop {
		_, _ = s.ChannelMessageSend(m.ChannelID, "It looks like you do not already have a recording running, did you mean to use the `"+config.DiscordPrefix+"start` command?")
		return
	}
	// Create new event
	lastEvent.Stop = time.Now()
	lastEvent.Orgasm = true
	lastEvent.Ejaculated = true
	err := service.CreateEvent(lastEvent)
	if err != nil {
		return
	}
	duration := lastEvent.Stop.Sub(lastEvent.Start)
	lastEvents := service.GetLastDayEventsForUser(user.ID)
	count := 0
	for _, event := range lastEvents {
		localTime := event.Start.Local()
		println(localTime.String())
		if localTime.Day() == time.Now().Day() {
			count++
		}
	}
	s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> just finished ("+strconv.Itoa(count)+" today)")
	var embed = discordgo.MessageEmbed{}
	embed.URL = "https://cad.bk1031.dev/events/" + lastEvent.ID
	// Embed description for DDD
	if lastEvent.Start.Local().Month() == 12 {
		embed.Description = strconv.Itoa(count) + " / Day " + strconv.Itoa(lastEvent.Start.Local().Day()) + " of DDD"
	}
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Started",
		Value:  lastEvent.Start.Local().Format("January 2 2006 3:4 pm"),
		Inline: true,
	})
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:   "Finished",
		Value:  lastEvent.Stop.Local().Format("January 2 2006 3:4 pm"),
		Inline: true,
	})
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  "Elapsed",
		Value: duration.String(),
	})
	embed.Title = "Summary"
	_, _ = service.Discord.ChannelMessageSendEmbed(m.ChannelID, &embed)
	// TODO: Add event to subscription queue
	return
}
