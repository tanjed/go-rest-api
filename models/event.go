package models

import (
	"time"

	"github.com/tanjed/go-rest-api/database"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	CreadtedBy  int       `json:"created_by"`
}

var events = []Event{}

func (event *Event) Save() {
	/*
	*DB will be introduced later
	 */
	events = append(events, *event)

}

func GetAllEvents() map[string]interface{} {
	return database.Get("SELECT * FROM events")
}
