package interfaces

import (
	"context"
	"github.com/google/uuid"
	"seatPlanner/internal/common/dto"
)

type PlanRepo interface {
	GetAll(ctx context.Context) ([]dto.Plan, error)
	GetById(ctx context.Context, uuid uuid.UUID) (dto.Plan, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
	Update(ctx context.Context, plan *dto.Plan, uuid uuid.UUID) error
	Insert(ctx context.Context, plan *dto.Plan) error
}

type SeatRepo interface {
	GetAll(ctx context.Context) ([]dto.Seat, error)
	GetById(ctx context.Context, uuid uuid.UUID) (dto.Seat, error)
	Delete(ctx context.Context, uuid uuid.UUID) error
	Update(ctx context.Context, plan *dto.Seat, uuid uuid.UUID) error
	Insert(ctx context.Context, plan *dto.Seat) error
}
