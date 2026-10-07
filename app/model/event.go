package model

import "time"

type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Venue     string    `json:"venue"`
	EventDate time.Time `json:"event_date"`
	Price     int       `json:"price"`
	Quota     int       `json:"quota"`
}

type CreateEventRequest struct {
	Name      string    `json:"name"`
	Venue     string    `json:"venue"`
	EventDate time.Time `json:"event_date"`
	Price     int       `json:"price"`
	Quota     int       `json:"quota"`
}

type ReplaceEventRequest struct {
	Name      string    `json:"name"`
	Venue     string    `json:"venue"`
	EventDate time.Time `json:"event_date"`
	Price     int       `json:"price"`
	Quota     int       `json:"quota"`
}
