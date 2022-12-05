package service

import (
	"server/model"
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
	result := DB.Where("id = ? OR discord_id = ?", userID).Find(&subscription)
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
