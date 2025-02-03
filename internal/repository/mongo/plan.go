package mongo

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	mongoDb "go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/net/context"
	"seatPlanner/internal/common/dto"
)

type PlanRepo struct {
	collection *mongoDb.Collection
}

func NewPlanRepo(db *mongoDb.Database) *PlanRepo {
	collections := db.Collection("plans")
	return &PlanRepo{collection: collections}
}

func (p *PlanRepo) GetAll(ctx context.Context) ([]dto.Plan, error) {
	cur, err := p.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var plans []dto.Plan

	for cur.Next(ctx) {

		var plan dto.Plan

		if err = cur.Decode(&plan); err != nil {
			return nil, err
		}

		plans = append(plans, plan)
	}

	return plans, nil
}

func (p *PlanRepo) GetById(ctx context.Context, uuid uuid.UUID) (dto.Plan, error) {
	cur := p.collection.FindOne(ctx, bson.M{"id": uuid})

	if cur.Err() != nil {
		return dto.Plan{}, cur.Err()
	}
	var plan dto.Plan
	if err := cur.Decode(&plan); err != nil {
		return dto.Plan{}, err
	}
	return plan, nil
}

func (p *PlanRepo) Delete(ctx context.Context, uuid uuid.UUID) error {
	_, err := p.collection.DeleteOne(ctx, bson.M{"id": uuid})

	if err != nil {
		return err
	}

	return nil
}

func (p *PlanRepo) Update(ctx context.Context, plan *dto.Plan, uuid uuid.UUID) error {
	filter := bson.M{"id": uuid}
	fields := bson.M{"$set": plan}

	_, err := p.collection.UpdateOne(ctx, filter, fields)

	if err != nil {
		return err
	}

	return nil
}

func (p *PlanRepo) Insert(ctx context.Context, plan *dto.Plan) error {
	_, err := p.collection.InsertOne(ctx, plan)

	if err != nil {
		return err
	}

	return nil
}
