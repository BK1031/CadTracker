package service

import (
	"server/model"
)

func GetAllEvents() []model.Event {
	var events []model.Event
	result := DB.Find(&events)
	if result.Error != nil {
	}
	return events
}

func GetEventByID(eventID string) model.Event {
	var event model.Event
	result := DB.Where("id = ?", eventID).Find(&event)
	if result.Error != nil {
	}
	return event
}

func CreateEvent(event model.Event) error {
	if DB.Where("id = ?", event.ID).Updates(&event).RowsAffected == 0 {
		println("New event created with id: " + event.ID)
		if result := DB.Create(&event); result.Error != nil {
			return result.Error
		}
	} else {
		println("Event with id: " + event.ID + " has been updated!")
	}
	return nil
}
