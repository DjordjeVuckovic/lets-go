package config

import (
	"fmt"
	"github.com/DjordjeVuckovic/lets-go/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	LogLevel    logger.Level
	LogHandler  logger.Handler
	Environment string
}

func LoadConfig() *Config {
	environment := os.Getenv("ENVIRONMENT")
	if err := godotenv.Load(); err != nil {
		if environment == "local" {
			panic(err)
		}
		fmt.Printf("Skipping .env file ...")
	}
	logLvl := os.Getenv("LOG_LEVEL")
	logHandler := os.Getenv("LOG_HANDLER")

	return &Config{
		LogLevel:    logger.Level(logLvl),
		LogHandler:  logger.Handler(logHandler),
		Environment: environment,
	}
}
