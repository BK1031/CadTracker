package model

import "time"

type Subscription struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	UserID         string    `json:"user_id"`
	DiscordGuild   string    `json:"discord_guild"`
	DiscordChannel string    `json:"discord_channel"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Subscription) TableName() string {
	return "subscription"
}
