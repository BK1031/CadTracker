package service

import (
	"github.com/bwmarrin/discordgo"
	"server/model"
	"strconv"
	"strings"
	"time"
)

func GetAllSubscriptions() []model.Subscription {
	var subscription []model.Subscription
	result := DB.Find(&subscription)
	if result.Error != nil {
	}
	return subscription
}

func GetAllSubscriptionsForUser(userID string) []model.Subscription {
	var subscription []model.Subscription
	result := DB.Where("user_id = ?", userID).Find(&subscription)
	if result.Error != nil {
	}
	return subscription
}

func GetAllSubscriptionsForChannel(channelID string) []model.Subscription {
	var subscription []model.Subscription
	result := DB.Where("discord_channel", channelID).Find(&subscription)
	if result.Error != nil {
	}
	return subscription
}

func GetSubscriptionByID(subscriptionID string) model.Subscription {
	var subscription model.Subscription
	result := DB.Where("id = ?", subscriptionID).Find(&subscription)
	if result.Error != nil {
	}
	return subscription
}

func CreateSubscription(subscription model.Subscription) error {
	if DB.Where("id = ?", subscription.ID).Updates(&subscription).RowsAffected == 0 {
		println("New subscription created with id: " + subscription.ID)
		if result := DB.Create(&subscription); result.Error != nil {
			return result.Error
		}
	} else {
		println("Subscription with id: " + subscription.ID + " has been updated!")
	}
	return nil
}

func DeleteSubscription(subscriptionID string) error {
	if result := DB.Where("id = ?", subscriptionID).Delete(model.Subscription{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func QueueSubscriptionEventForUser(user model.User, event model.Event, start bool) {
	subscriptions := GetAllSubscriptionsForUser(user.ID)
	println("Sending " + strconv.Itoa(len(subscriptions)) + " event updates for " + user.FirstName + " (<@" + user.DiscordID + ">)")
	for _, subscription := range subscriptions {
		if start {
			// Start Event
			Discord.ChannelMessageSend(subscription.DiscordChannel, user.FirstName+" (<@"+user.DiscordID+">) just started")
		} else {
			// Stop Event
			duration := event.Stop.Sub(event.Start)
			// Find which number this event is for the current day
			lastEvents := GetLastDayEventsForUser(user.ID)
			count := 0
			for _, e := range lastEvents {
				localTime := e.Start.Local()
				if localTime.Day() == time.Now().Day() {
					count++
				}
			}
			Discord.ChannelMessageSend(subscription.DiscordChannel, user.FirstName+" (<@"+user.DiscordID+">) just finished ("+strconv.Itoa(count)+" today)")
			// Create the summary embed
			var embed = discordgo.MessageEmbed{}
			embed.URL = "https://cad.bk1031.dev/events/" + event.ID
			// Embed description for DDD
			if event.Start.Local().Month() == 12 {
				embed.Description = strconv.Itoa(count) + " / Day " + strconv.Itoa(event.Start.Local().Day()) + " of DDD"
			}
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Started",
				Value:  event.Start.Local().Format("January 2, 2006 3:04 pm"),
				Inline: true,
			})
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "Finished",
				Value:  event.Stop.Local().Format("January 2, 2006 3:04 pm"),
				Inline: true,
			})
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:  "Elapsed",
				Value: strings.Replace(duration.String(), "m", "m ", 1),
			})
			embed.Title = "Summary"
			_, _ = Discord.ChannelMessageSendEmbed(subscription.DiscordChannel, &embed)
		}
		println("Send update to Guild [" + subscription.DiscordGuild + "] - Channel [" + subscription.DiscordChannel + "]")
	}
	return
}
