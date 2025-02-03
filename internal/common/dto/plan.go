package dto

import "github.com/google/uuid"

type Plan struct {
	Id         uuid.UUID `json:"id"`
	Seat       []Seat    `json:"seats"`
	Background string    `json:"background"`
}
