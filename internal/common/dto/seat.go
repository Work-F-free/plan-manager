package dto

import "github.com/google/uuid"

type Seat struct {
	Id       uuid.UUID `json:"id"`
	Color    string    `json:"color"`
	SeatType string    `json:"type"`
	CoordX   float64   `json:"coord_x"`
	CoordY   float64   `json:"coord_y"`
}
