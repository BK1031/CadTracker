package service

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"server/config"
	"server/model"
	"time"
)

var DB *gorm.DB

var dbRetries = 0

func InitializeDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=cad_tracker port=%s sslmode=disable TimeZone=UTC", config.PostgresHost, config.PostgresUser, config.PostgresPassword, config.PostgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if dbRetries < 15 {
			dbRetries++
			println("failed to connect database, retrying in 5s... ")
			time.Sleep(time.Second * 5)
			InitializeDB()
		} else {
			println("failed to connect database after 15 attempts, terminating program...")
			os.Exit(100)
		}
	} else {
		println("Connected to postgres database")
		db.AutoMigrate(&model.User{}, &model.Privacy{}, &model.Event{}, &model.Subscription{})
		println("AutoMigration complete")
		DB = db
	}
}
