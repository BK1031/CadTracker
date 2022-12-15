package controller

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"server/config"
	"server/service"
	"strings"
)

func InitializeDiscordBot() {
	service.Discord.AddHandler(OnDiscordMessage)
	service.Discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err := service.Discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Discord Bot is now running! [Prefix = " + config.DiscordPrefix + "]")
}

func OnDiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// or messages that don't start with the prefix
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, config.DiscordPrefix) {
		return
	}
	if strings.HasPrefix(strings.Split(m.Content, config.DiscordPrefix)[1], "ping") {
		DiscordPing(s, m)
	}
	if strings.HasPrefix(strings.Split(m.Content, config.DiscordPrefix)[1], "invite") {
		DiscordInvite(s, m)
	}
	if strings.HasPrefix(strings.Split(m.Content, config.DiscordPrefix)[1], "link") {
		DiscordLinkAccount(s, m)
	}
	if strings.HasPrefix(strings.Split(m.Content, config.DiscordPrefix)[1], "start") {
		DiscordStartEvent(s, m)
	}
	if strings.HasPrefix(strings.Split(m.Content, config.DiscordPrefix)[1], "stop") {
		DiscordStopEvent(s, m)
	}
	if strings.HasPrefix(strings.Split(m.Content, config.DiscordPrefix)[1], "sub") {
		DiscordCreateSubscription(s, m)
	}
	if strings.HasPrefix(strings.Split(m.Content, config.DiscordPrefix)[1], "unsub") {
		DiscordRemoveSubscription(s, m)
	}
}
