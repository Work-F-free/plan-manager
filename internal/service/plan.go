package service

import (
	"github.com/google/uuid"
	"seatPlanner/internal/common/dto"
	"seatPlanner/internal/repository"
	"seatPlanner/internal/repository/interfaces"
)

type Plan struct {
	seatRepo *interfaces.SeatRepo
	planRepo *interfaces.PlanRepo
}

func NewPlan(repo *repository.Repo) *Plan {
	return &Plan{
		seatRepo: &repo.SeatRepo,
		planRepo: &repo.PlanRepo,
	}
}

func (p Plan) GetSeat(uuid uuid.UUID) (dto.Seat, int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) GetAllSeats() ([]dto.Seat, int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) CreateSeat(seat dto.Seat) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) DeleteSeat(uuid uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) UpdateSeat(seat dto.Seat, uuid uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) GetPlan(uuid uuid.UUID) (dto.Plan, int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) GetAllPlans() ([]dto.Plan, int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) CreatePlan(seat dto.Plan) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) DeletePlan(uuid uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p Plan) UpdatePlan(seat dto.Plan, uuid uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}
