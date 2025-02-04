package dto

import "github.com/google/uuid"

type Seat struct {
	Id       uuid.UUID `json:"id"`
	Color    string    `json:"color"`
	SeatType string    `json:"type"`
	SeatNum  string    `json:"number_seat"`
	CoordX   float64   `json:"coord_x"`
	CoordY   float64   `json:"coord_y"`
}
