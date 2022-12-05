package model

import "time"

type Event struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	UserID     string    `json:"user_id"`
	Start      time.Time `json:"start"`
	Stop       time.Time `json:"stop"`
	Notes      string    `json:"notes"`
	Orgasm     bool      `json:"orgasm"`
	Ejaculated bool      `json:"ejaculated"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Event) TableName() string {
	return "event"
}
