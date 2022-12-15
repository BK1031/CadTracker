package controller

import "github.com/bwmarrin/discordgo"

func DiscordInvite(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "https://discord.com/oauth2/authorize?client_id=1049305686435176568&permissions=68672&scope=bot")
}
