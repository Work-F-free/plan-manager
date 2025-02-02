package dto

import "github.com/google/uuid"

type Seat struct {
	id        uuid.UUID
	color     string
	seatTyper string
	coordX    float64
	coordY    float64
}
