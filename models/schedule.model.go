package models

import "time"

type Schedule struct {
	ID          int       `json:"id"`
	CityID      int       `json:"city_id"`
	CinemaID    int       `json:"cinema_id"`
	MovieID     int       `json:"movie_id"`
	StartAt     string    `json:"start_at"` // format "15:04"
	TicketPrice int       `json:"ticket_price"`
	ShowDate    string    `json:"show_date"` // format "2006-01-02"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
