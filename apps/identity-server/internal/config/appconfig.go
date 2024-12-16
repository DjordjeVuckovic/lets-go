package config

import (
	"fmt"
	"github.com/DjordjeVuckovic/lets-go/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	LogLevel    logger.Level
	LogHandler  logger.Handler
	Environment string
}

func LoadAppConfig() *AppConfig {
	env := os.Getenv("ENVIRONMENT")
	if err := godotenv.Load(); err != nil {
		if env == "local" {
			panic(err)
		}
		fmt.Printf("Skipping .env file ...")
	}
	logLvl := os.Getenv("LOG_LEVEL")
	logHandler := os.Getenv("LOG_HANDLER")

	return &AppConfig{
		LogLevel:    logger.Level(logLvl),
		LogHandler:  logger.Handler(logHandler),
		Environment: env,
	}
}
