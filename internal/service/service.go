package service

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"seatPlanner/internal/common/dto"
	"seatPlanner/internal/repository"
)

type PlannerService interface {
	GetSeat(ctx context.Context, uuid uuid.UUID) (dto.Seat, int, error)
	GetAllSeats(ctx context.Context) ([]dto.Seat, int, error)
	CreateSeat(ctx context.Context, seat dto.Seat) (int, error)
	DeleteSeat(ctx context.Context, uuid uuid.UUID) (int, error)
	UpdateSeat(ctx context.Context, seat dto.Seat, uuid uuid.UUID) (int, error)

	GetPlan(ctx context.Context, uuid uuid.UUID) (dto.Plan, int, error)
	GetAllPlans(ctx context.Context) ([]dto.Plan, int, error)
	CreatePlan(ctx context.Context, seat dto.Plan) (int, error)
	DeletePlan(ctx context.Context, uuid uuid.UUID) (int, error)
	UpdatePlan(ctx context.Context, seat dto.Plan, uuid uuid.UUID) (int, error)
}

type Service struct {
	PlannerService
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		PlannerService: NewPlan(repo),
	}
}
