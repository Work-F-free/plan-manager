package interfaces

import (
	"context"
	"github.com/google/uuid"
	"seatPlanner/internal/common/dto"
)

type PlanRepo interface {
	GetAll(ctx context.Context) ([]dto.Plan, error)
	GetById(ctx context.Context, uuid uuid.UUID) (dto.Plan, error)
	Delete(uuid uuid.UUID) error
	Update(plan *dto.Plan, uuid uuid.UUID) error
	Insert(plan *dto.Plan) error
}

type SeatRepo interface {
	GetAll(ctx context.Context) ([]dto.Seat, error)
	GetById(ctx context.Context, uuid uuid.UUID) (dto.Seat, error)
	Delete(uuid uuid.UUID) error
	Update(plan *dto.Seat, uuid uuid.UUID) error
	Insert(plan *dto.Seat) error
}
