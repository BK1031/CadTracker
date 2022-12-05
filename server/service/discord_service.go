package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"server/config"
)

var Discord *discordgo.Session

func ConnectDiscord() {
	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		fmt.Println("Error creating Discord session, ", err)
		return
	}
	Discord = dg
}
