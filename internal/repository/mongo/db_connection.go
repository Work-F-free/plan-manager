package mongo

import (
	"fmt"
	mongoDb "go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/net/context"
	"seatPlanner/internal/common/config"
	"seatPlanner/internal/common/connection"
)

type Connection struct {
	client *mongoDb.Client
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect(config config.DBConfig, ctx context.Context) (connection.Connection, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", config.DBHost, config.DBPort)

	client, err := mongoDb.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return connection.Connection{}, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return connection.Connection{}, err
	}

	db := client.Database(config.DBName)

	return connection.Connection{Database: db}, nil
}

func (c *Connection) Disconnect(ctx context.Context) error {
	err := c.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}
