package mongo

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	mongoDb "go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/net/context"
	"seatPlanner/internal/common/dto"
)

type SeatRepo struct {
	collection *mongoDb.Collection
}

func NewSeatRepo(db *mongoDb.Database) *SeatRepo {
	collections := db.Collection("seats")
	return &SeatRepo{
		collection: collections,
	}
}

func (s *SeatRepo) GetAll(ctx context.Context) ([]dto.Seat, error) {
	cur, err := s.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var plans []dto.Seat

	for cur.Next(ctx) {

		var seat dto.Seat

		if err = cur.Decode(&seat); err != nil {
			return nil, err
		}

		plans = append(plans, seat)
	}

	return plans, nil
}

func (s *SeatRepo) GetById(ctx context.Context, uuid uuid.UUID) (dto.Seat, error) {
	cur := s.collection.FindOne(ctx, bson.M{"id": uuid})

	if cur.Err() != nil {
		return dto.Seat{}, cur.Err()
	}
	var seat dto.Seat
	if err := cur.Decode(&seat); err != nil {
		return dto.Seat{}, err
	}
	return seat, nil
}

func (s *SeatRepo) Delete(ctx context.Context, uuid uuid.UUID) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": uuid})

	if err != nil {
		return err
	}

	return nil
}

func (s *SeatRepo) Update(ctx context.Context, seat *dto.Seat, uuid uuid.UUID) error {
	filter := bson.M{"id": uuid}
	fields := bson.M{"$set": seat}

	_, err := s.collection.UpdateOne(ctx, filter, fields)

	if err != nil {
		return err
	}

	return nil
}

func (s *SeatRepo) Insert(ctx context.Context, seat *dto.Seat) error {
	_, err := s.collection.InsertOne(ctx, seat)

	if err != nil {
		return err
	}

	return nil
}
