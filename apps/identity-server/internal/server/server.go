package server

import (
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/internal/db"
	mymiddleware "github.com/DjordjeVuckovic/lets-go/pkg/api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	GracefulShutdownTimeout = 10 * time.Second
)

type Server struct {
	E  *echo.Echo
	ec *EnvConfig
}

func NewServer(e *echo.Echo, ec *EnvConfig, database *db.Database) *Server {
	e.DisableHTTP2 = !ec.UseHttp2

	s := &Server{
		E: e, ec: ec,
	}

	s.setupMiddlewares()
	s.setupHealthChecks(database)

	return s
}

func (s *Server) setupMiddlewares() {
	s.E.Use(mymiddleware.Logger())
	s.E.Use(middleware.Recover())
	s.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: s.ec.CorsOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}

func (s *Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := s.E.Start(":" + s.ec.Port); err != nil && err != http.ErrServerClosed {
			s.E.Logger.Fatal("shutting down the server")
		}
	}()
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), GracefulShutdownTimeout)
	defer cancel()
	if err := s.E.Shutdown(ctx); err != nil {
		s.E.Logger.Fatal(err)
		return err
	}
	return nil
}

func (s *Server) setupHealthChecks(db *db.Database) {
	s.E.GET("/health", func(c echo.Context) error {
		err := db.DB.Ping()
		if err != nil {
			return err
		}
		return c.JSON(200, "ok")
	})
	s.E.GET("/ready", func(c echo.Context) error {
		return c.JSON(200, "ok")
	})
}
