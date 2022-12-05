package model

import "time"

type Privacy struct {
	UserID        string    `gorm:"primaryKey" json:"user_id"`
	Status        string    `json:"status"`
	StatsBasic    string    `json:"stats_basic"`
	StatsDetailed string    `json:"stats_detailed"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Privacy) TableName() string {
	return "user_privacy"
}
