package repository

import (
	mongoDb "go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/net/context"
	"seatPlanner/internal/common/config"
	"seatPlanner/internal/repository/interfaces"
	"seatPlanner/internal/repository/mongo"
)

type Connection struct {
	*mongoDb.Database
}

type DBConnection interface {
	Connect(config config.DBConfig, ctx context.Context) (Connection, error)
	Disconnect(ctx context.Context) error
}

type Repo struct {
	interfaces.SeatRepo
	interfaces.PlanRepo
}

func New(cnt Connection) *Repo {
	return &Repo{
		PlanRepo: mongo.NewPlanRepo(cnt.Database),
		SeatRepo: mongo.NewSeatRepo(cnt.Database),
	}
}
