package service

import (
	"server/model"
)

func GetAllUsers() []model.User {
	var users []model.User
	result := DB.Find(&users)
	if result.Error != nil {
	}
	for i := range users {
		users[i].Privacy = GetPrivacyForUser(users[i].ID)
	}
	return users
}

func GetUserByID(userID string) model.User {
	var user model.User
	result := DB.Where("id = ? OR discord_id = ?", userID, userID).Find(&user)
	if result.Error != nil {
	}
	user.Privacy = GetPrivacyForUser(user.ID)
	return user
}

func CreateUser(user model.User) error {
	if DB.Where("id = ?", user.ID).Updates(&user).RowsAffected == 0 {
		println("New user created with id: " + user.ID)
		if result := DB.Create(&user); result.Error != nil {
			return result.Error
		}
	} else {
		println("User with id: " + user.ID + " has been updated!")
	}
	if user.Privacy.UserID != "" {
		println("User with id: " + user.ID + " has non-empty privacy object, setting privacy in db...")
		user.Privacy.UserID = user.ID
		if err := SetPrivacyForUser(user.ID, user.Privacy); err != nil {
			return err
		}
	} else {
		println("User with id: " + user.ID + " has empty privacy object, nothing to do here!")
	}
	return nil
}
