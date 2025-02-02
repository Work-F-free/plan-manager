package service

import (
	"github.com/google/uuid"
	"seatPlanner/internal/common/dto"
	"seatPlanner/internal/repository"
)

type PlannerService interface {
	GetSeat(uuid uuid.UUID) (dto.Seat, int, error)
	GetAllSeats() ([]dto.Seat, int, error)
	CreateSeat(seat dto.Seat) (int, error)
	DeleteSeat(uuid uuid.UUID) (int, error)
	UpdateSeat(seat dto.Seat, uuid uuid.UUID) (int, error)

	GetPlan(uuid uuid.UUID) (dto.Plan, int, error)
	GetAllPlans() ([]dto.Plan, int, error)
	CreatePlan(seat dto.Plan) (int, error)
	DeletePlan(uuid uuid.UUID) (int, error)
	UpdatePlan(seat dto.Plan, uuid uuid.UUID) (int, error)
}

type Service struct {
	PlannerService
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		PlannerService: NewPlan(repo),
	}
}
