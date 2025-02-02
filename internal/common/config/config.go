package config

import "os"

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBPassword string
	DBUsername string
}

type MinIOConfig struct {
	MinIOUsername   string
	MinIOPassword   string
	MinioEndpoint   string
	MinIOBucketName string
	MinIOUseSSL     bool
}

type Config struct {
	Port string
	DBConfig
	MinIOConfig
}

var AppConfig *Config

func LoadConfig() {

	dbCfg := DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUsername: os.Getenv("DB_USERNAME"),
	}

	minCfg := MinIOConfig{
		MinIOUsername:   os.Getenv("MINIO_USERNAME"),
		MinIOPassword:   os.Getenv("MINIO_PASSWORD"),
		MinIOBucketName: os.Getenv("MINIO_BUCKET_NAME"),
		MinioEndpoint:   os.Getenv("MINIO_ENDPOINT"),
		MinIOUseSSL:     os.Getenv("MINIO_USE_SSL") == "true",
	}

	AppConfig = &Config{
		Port:        os.Getenv("PORT"),
		DBConfig:    dbCfg,
		MinIOConfig: minCfg,
	}
}
