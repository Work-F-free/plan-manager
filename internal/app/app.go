package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"seatPlanner/internal/common/config"
	"seatPlanner/internal/handler"
	"seatPlanner/internal/repository"
	"seatPlanner/internal/repository/mongo"
	"seatPlanner/internal/server"
	"seatPlanner/internal/service"
)

func Run() error {
	config.LoadConfig()
	ctx := context.Background()
	serv := new(server.Server)

	connection, err := mongo.NewConnection().Connect(config.AppConfig.DBConfig, ctx)
	if err != nil {
		logrus.Fatalf("error occured while сonnecting DB: %s", err.Error())
		return err
	}

	repo := repository.New(connection)

	planService := service.NewService(repo)

	handlers := handler.New(planService)
	if err = serv.Run(os.Getenv("API_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running Http-Server: %s", err.Error())
		return err
	}
	return nil
}
