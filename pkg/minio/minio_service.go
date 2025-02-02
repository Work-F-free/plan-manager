package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"seatPlanner/internal/common/config"
	"seatPlanner/pkg/minio/helpers"
	"sync"
	"time"
)

func (m *minioClient) CreateOne(file helpers.FileDataType) (string, error) {
	objectID := uuid.New().String()

	reader := bytes.NewReader(file.Data)

	_, err := m.mc.PutObject(context.Background(), config.AppConfig.MinIOBucketName, objectID, reader, int64(len(file.Data)), minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("ошибка при создании объекта %s: %v", file.FileName, err)
	}

	url, err := m.mc.PresignedGetObject(context.Background(), config.AppConfig.MinIOBucketName, objectID, time.Second*24*60*60, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка при создании URL для объекта %s: %v", file.FileName, err)
	}

	return url.String(), nil
}

func (m *minioClient) CreateMany(data map[string]helpers.FileDataType) ([]string, error) {
	urls := make([]string, 0, len(data))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urlCh := make(chan string, len(data))

	var wg sync.WaitGroup

	for objectID, file := range data {
		wg.Add(1)
		go func(objectID string, file helpers.FileDataType) {
			defer wg.Done()                                                                                                                                        // Уменьшение счетчика WaitGroup после завершения горутины.
			_, err := m.mc.PutObject(ctx, config.AppConfig.MinIOBucketName, objectID, bytes.NewReader(file.Data), int64(len(file.Data)), minio.PutObjectOptions{}) // Создание объекта в бакете MinIO.
			if err != nil {
				cancel()
				return
			}

			url, err := m.mc.PresignedGetObject(ctx, config.AppConfig.MinIOBucketName, objectID, time.Second*24*60*60, nil)
			if err != nil {
				cancel()
				return
			}

			urlCh <- url.String()
		}(objectID, file)
	}

	go func() {
		wg.Wait()
		close(urlCh)
	}()

	for url := range urlCh {
		urls = append(urls, url)
	}

	return urls, nil
}

func (m *minioClient) GetOne(objectID string) (string, error) {
	url, err := m.mc.PresignedGetObject(context.Background(), config.AppConfig.MinIOBucketName, objectID, time.Second*24*60*60, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка при получении URL для объекта %s: %v", objectID, err)
	}

	return url.String(), nil
}

func (m *minioClient) GetMany(objectIDs []string) ([]string, error) {
	urlCh := make(chan string, len(objectIDs))
	errCh := make(chan helpers.OperationError, len(objectIDs))

	var wg sync.WaitGroup
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, objectID := range objectIDs {
		wg.Add(1)
		go func(objectID string) {
			defer wg.Done()
			url, err := m.GetOne(objectID)
			if err != nil {
				errCh <- helpers.OperationError{ObjectID: objectID, Error: fmt.Errorf("ошибка при получении объекта %s: %v", objectID, err)}
				cancel()
				return
			}
			urlCh <- url
		}(objectID)
	}

	go func() {
		wg.Wait()
		close(urlCh)
		close(errCh)
	}()

	var urls []string
	var errs []error
	for url := range urlCh {
		urls = append(urls, url)
	}
	for opErr := range errCh {
		errs = append(errs, opErr.Error)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("ошибки при получении объектов: %v", errs) // Возврат ошибки, если возникли ошибки при получении объектов
	}

	return urls, nil
}

func (m *minioClient) DeleteOne(objectID string) error {
	err := m.mc.RemoveObject(context.Background(), config.AppConfig.MinIOBucketName, objectID, minio.RemoveObjectOptions{})
	if err != nil {
		return err // Возвращаем ошибку, если не удалось удалить объект.
	}
	return nil
}

func (m *minioClient) DeleteMany(objectIDs []string) error {
	errCh := make(chan helpers.OperationError, len(objectIDs))
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, objectID := range objectIDs {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			err := m.mc.RemoveObject(ctx, config.AppConfig.MinIOBucketName, id, minio.RemoveObjectOptions{})
			if err != nil {
				errCh <- helpers.OperationError{ObjectID: id, Error: fmt.Errorf("ошибка при удалении объекта %s: %v", id, err)}
				cancel()
			}
		}(objectID)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	var errs []error
	for opErr := range errCh {
		errs = append(errs, opErr.Error)
	}

	if len(errs) > 0 {
		return fmt.Errorf("ошибки при удалении объектов: %v", errs) // Возврат ошибки, если возникли ошибки при удалении объектов
	}

	return nil
}
