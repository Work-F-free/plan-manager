package connection

import mongoDb "go.mongodb.org/mongo-driver/v2/mongo"

type Connection struct {
	*mongoDb.Database
}
