package models

import (
	"errors"
	"time"

	"github.com/tanjed/go-rest-api/database"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	CreadtedBy  int       `json:"created_by"`
}

func (event *Event) Save() error {
	query := `
INSERT INTO events (name, description, location, date_time, created_by)
VALUES (?, ?, ?, ?, ?)`

	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.CreadtedBy)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	event.ID = id
	return err
}
func GetById(id string) (*Event, error) {
	query := `
SELECT * FROM events WHERE id = ? LIMIT 1
	`
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()
	var event Event
	var dateTime []uint8
	err = statement.QueryRow(id).Scan(&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&dateTime,
		&event.CreadtedBy)

	if err != nil {
		return nil, err
	}

	datetimeStr := string([]byte(dateTime))

	t, err := time.Parse("2006-01-02 15:04:05", datetimeStr)
	if err != nil {
		return nil, err
	}
	event.DateTime = t

	return &event, nil
}

func (event *Event) Update() (bool, error) {
	query := `
UPDATE events SET
name = ?,
description = ?,
location = ?,
date_time = ?
WHERE id = ?
`

	statement, err := database.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	result, err := statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	if err != nil {
		return false, err
	}

	effectedRows, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	if effectedRows <= 0 {
		return false, errors.New("no updatable entry found")
	}

	return true, nil
}
func (event *Event) Delete() (bool, error) {
	query := `
DELETE FROM events
WHERE id = ?
	`

	statement, err := database.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	result, err := statement.Exec(event.ID)

	if err != nil {
		return false, err
	}

	effectedRows, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	if effectedRows <= 0 {
		return false, errors.New("no deletable entry found")
	}

	return true, nil
}

func GetAllEvents() (*[]Event, error) {
	rows, err := database.DB.Query("SELECT * from events")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		var dateTime []uint8
		err := rows.Scan(&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&dateTime,
			&event.CreadtedBy)

		if err != nil {
			return nil, err
		}

		datetimeStr := string([]byte(dateTime))

		t, err := time.Parse("2006-01-02 15:04:05", datetimeStr)
		if err != nil {
			return nil, err
		}
		event.DateTime = t
		events = append(events, event)
	}

	return &events, nil

}

func New(name, description, location string, dateTime time.Time) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
	}
}
