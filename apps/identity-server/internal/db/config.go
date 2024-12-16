package db

import (
	"fmt"
	"os"
)

const (
	ConnectionPoolSize    = 25
	ConnectionMaxIdleTime = 25
	ConnectionMaxLifetime = 5 * 60 * 60
)

type Config struct {
	ConnectionString         string
	DbName                   string
	ConnectionPoolSize       int
	ConnectionMaxLifetimeSec int
	ConnectionMaxIdleTime    int
}

func LoadConfig() (*Config, error) {
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, fmt.Errorf("DB_PORT is not set")
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("DB_HOST is not set")
	}
	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		return nil, fmt.Errorf("DB_USERNAME is not set")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD is not set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB_NAME is not set")
	}
	dbSslMode := os.Getenv("DB_SSL_MODE")
	if dbSslMode == "" {
		dbSslMode = "disable"
	}
	connString := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSslMode
	return &Config{
		ConnectionString:         connString,
		DbName:                   dbName,
		ConnectionPoolSize:       ConnectionPoolSize,
		ConnectionMaxIdleTime:    ConnectionMaxIdleTime,
		ConnectionMaxLifetimeSec: ConnectionMaxLifetime,
	}, nil
}
