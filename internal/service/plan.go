package service

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"net/http"
	"seatPlanner/internal/common/dto"
	"seatPlanner/internal/repository"
	"seatPlanner/internal/repository/interfaces"
)

type Plan struct {
	seatRepo interfaces.SeatRepo
	planRepo interfaces.PlanRepo
}

func NewPlan(repo *repository.Repo) *Plan {
	return &Plan{
		seatRepo: repo.SeatRepo,
		planRepo: repo.PlanRepo,
	}
}

func (p *Plan) GetSeat(ctx context.Context, uuid uuid.UUID) (dto.Seat, int, error) {
	seat, err := p.seatRepo.GetById(ctx, uuid)
	if err != nil {
		return dto.Seat{}, http.StatusInternalServerError, err
	}

	return seat, http.StatusOK, nil
}

func (p *Plan) GetAllSeats(ctx context.Context) ([]dto.Seat, int, error) {
	seats, err := p.seatRepo.GetAll(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return seats, http.StatusOK, nil
}

func (p *Plan) CreateSeat(ctx context.Context, seat *dto.Seat) (int, error) {
	err := p.seatRepo.Insert(ctx, seat)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (p *Plan) DeleteSeat(ctx context.Context, uuid uuid.UUID) (int, error) {
	err := p.seatRepo.Delete(ctx, uuid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

func (p *Plan) UpdateSeat(ctx context.Context, seat dto.Seat, uuid uuid.UUID) (int, error) {
	err := p.seatRepo.Update(ctx, &seat, uuid)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

func (p *Plan) GetPlan(ctx context.Context, uuid uuid.UUID) (dto.Plan, int, error) {
	plan, err := p.planRepo.GetById(ctx, uuid)
	if err != nil {
		return dto.Plan{}, http.StatusInternalServerError, err
	}
	return plan, http.StatusOK, nil
}

func (p *Plan) GetAllPlans(ctx context.Context) ([]dto.Plan, int, error) {
	plans, err := p.planRepo.GetAll(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return plans, http.StatusOK, nil
}

func (p *Plan) CreatePlan(ctx context.Context, seat dto.Plan) (int, error) {
	err := p.planRepo.Insert(ctx, &seat)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (p *Plan) DeletePlan(ctx context.Context, uuid uuid.UUID) (int, error) {
	err := p.planRepo.Delete(ctx, uuid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}

func (p *Plan) UpdatePlan(ctx context.Context, seat dto.Plan, uuid uuid.UUID) (int, error) {
	err := p.planRepo.Update(ctx, &seat, uuid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
