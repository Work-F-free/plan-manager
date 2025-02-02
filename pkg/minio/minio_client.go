package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"seatPlanner/internal/common/config"
	"seatPlanner/pkg/minio/helpers"
)

type Client interface {
	InitMinio() error
	CreateOne(file helpers.FileDataType) (string, error)
	CreateMany(map[string]helpers.FileDataType) ([]string, error)
	GetOne(objectID string) (string, error)
	GetMany(objectIDs []string) ([]string, error)
	DeleteOne(objectID string) error
	DeleteMany(objectIDs []string) error
}

type minioClient struct {
	mc *minio.Client
}

func NewMinioClient() Client {
	return &minioClient{}
}

func (m *minioClient) InitMinio() error {
	ctx := context.Background()

	client, err := minio.New(config.AppConfig.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AppConfig.MinIOUsername, config.AppConfig.MinIOPassword, ""),
		Secure: config.AppConfig.MinIOUseSSL,
	})
	if err != nil {
		return err
	}

	m.mc = client

	exists, err := m.mc.BucketExists(ctx, config.AppConfig.MinIOBucketName)
	if err != nil {
		return err
	}
	if !exists {
		err := m.mc.MakeBucket(ctx, config.AppConfig.MinIOBucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
