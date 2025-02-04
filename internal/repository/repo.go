package repository

import (
	"golang.org/x/net/context"
	"seatPlanner/internal/common/config"
	"seatPlanner/internal/common/connection"
	"seatPlanner/internal/repository/interfaces"
	"seatPlanner/internal/repository/mongo"
)

type DBConnection interface {
	Connect(config config.DBConfig, ctx context.Context) (connection.Connection, error)
	Disconnect(ctx context.Context) error
}

type Repo struct {
	interfaces.SeatRepo
	interfaces.PlanRepo
}

func New(cnt connection.Connection) *Repo {
	return &Repo{
		PlanRepo: mongo.NewPlanRepo(cnt.Database),
		SeatRepo: mongo.NewSeatRepo(cnt.Database),
	}
}
