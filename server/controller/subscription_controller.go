package controller

import (
	"github.com/bwmarrin/discordgo"
	"server/config"
	"server/model"
	"server/service"
	"strings"
	"time"
)

func DiscordCreateSubscription(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(strings.Split(m.Content, " ")) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Command usage: `"+config.DiscordPrefix+"sub [CadTracker Account ID / @mention]`")
		return
	}
	var subUser model.User
	if len(m.Mentions) > 0 && m.Mentions[0].ID != "" {
		subUser = service.GetUserByID(m.Mentions[0].ID)
	} else {
		subUser = service.GetUserByID(strings.Split(m.Content, config.DiscordPrefix+"sub ")[1])
	}
	if subUser.ID == "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "No CadTracker account found with that ID!")
		return
	}
	subscription := model.Subscription{
		ID:             m.ChannelID + "-" + subUser.ID,
		UserID:         subUser.ID,
		DiscordGuild:   m.GuildID,
		DiscordChannel: m.ChannelID,
		UpdatedAt:      time.Time{},
		CreatedAt:      time.Time{},
	}
	err := service.CreateSubscription(subscription)
	if err != nil {
		return
	}
	if len(m.Mentions) > 0 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Successfully subscribed this channel to <@"+subUser.DiscordID+"> on CadTracker!")
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Successfully subscribed this channel to "+subUser.FirstName+" "+subUser.LastName+" on CadTracker!")
	}
}

func DiscordRemoveSubscription(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(strings.Split(m.Content, " ")) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Command usage: `"+config.DiscordPrefix+"unsub [CadTracker Account ID / @mention]`")
		return
	}
	var subUser model.User
	if len(m.Mentions) > 0 && m.Mentions[0].ID != "" {
		subUser = service.GetUserByID(m.Mentions[0].ID)
	} else {
		subUser = service.GetUserByID(strings.Split(m.Content, config.DiscordPrefix+"unsub ")[1])
	}
	if subUser.ID == "" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "No CadTracker account found with that ID!")
		return
	}
	err := service.DeleteSubscription(m.ChannelID + "-" + subUser.ID)
	if err != nil {
		return
	}
	if len(m.Mentions) > 0 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "This channel will no longer receive notifications from <@"+subUser.DiscordID+"> on CadTracker!")
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "This channel will no longer receive notifications from "+subUser.FirstName+" "+subUser.LastName+" on CadTracker!")
	}
}
