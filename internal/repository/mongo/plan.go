package mongo

import (
	"github.com/google/uuid"
	mongoDb "go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/net/context"
	"seatPlanner/internal/common/dto"
)

type PlanRepo struct {
	client *mongoDb.Client
}

func NewPlanRepo(client *mongoDb.Client) *PlanRepo {
	return &PlanRepo{client: client}
}

func (p PlanRepo) GetAll(ctx context.Context) ([]dto.Plan, error) {
	//TODO implement me
	panic("implement me")
}

func (p PlanRepo) GetById(ctx context.Context, uuid uuid.UUID) (dto.Plan, error) {
	//TODO implement me
	panic("implement me")
}

func (p PlanRepo) Delete(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p PlanRepo) Update(plan *dto.Plan, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p PlanRepo) Insert(plan *dto.Plan) error {
	//TODO implement me
	panic("implement me")
}
