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

	AppConfig = &Config{
		Port:     os.Getenv("PORT"),
		DBConfig: dbCfg,
	}
}
