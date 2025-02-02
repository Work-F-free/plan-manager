package dto

import "github.com/google/uuid"

type Plan struct {
	id         uuid.UUID
	seat       []Seat
	background string
}
