package model

import "time"

type User struct {
	ID                string    `gorm:"primaryKey" json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `gorm:"unique" json:"email"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	Gender            string    `json:"gender"`
	DiscordID         string    `json:"discord_id"`
	Privacy           Privacy   `gorm:"-" json:"privacy"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (User) TableName() string {
	return "user"
}

func (user User) String() string {
	return "(" + user.ID + ")" + " " + user.FirstName + " " + user.LastName
}
