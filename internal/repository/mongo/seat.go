package mongo

import (
	"github.com/google/uuid"
	mongoDb "go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/net/context"
	"seatPlanner/internal/common/dto"
)

type SeatRepo struct {
	client *mongoDb.Client
}

func NewSeatRepo(client *mongoDb.Client) *SeatRepo {
	return &SeatRepo{
		client: client,
	}
}

func (s SeatRepo) GetAll(ctx context.Context) ([]dto.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (s SeatRepo) GetById(ctx context.Context, uuid uuid.UUID) (dto.Seat, error) {
	//TODO implement me
	panic("implement me")
}

func (s SeatRepo) Delete(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s SeatRepo) Update(plan *dto.Seat, uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s SeatRepo) Insert(plan *dto.Seat) error {
	//TODO implement me
	panic("implement me")
}
