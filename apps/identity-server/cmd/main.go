package main

import (
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/api"
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/internal/config"
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/internal/db"
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/internal/server"
	"github.com/DjordjeVuckovic/lets-go/pkg/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	appConfig := config.LoadConfig()
	logConfig := logger.Config{
		Level:   appConfig.LogLevel,
		Handler: appConfig.LogHandler,
	}
	logger.Init(logConfig)

	dbConfig, err := db.LoadConfig()
	if err != nil {
		panic(err)
	}
	database, dbErr := db.NewDatabase(dbConfig)
	if dbErr != nil {
		panic(dbErr)
	}

	serverConfig, err := server.LoadConfig()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	s := server.NewServer(e, serverConfig, database)

	a := api.NewAuthRouter(e, database)
	a.RegisterRoutes()

	if err := s.Start(); err != nil {
		panic(err)
	}

}
