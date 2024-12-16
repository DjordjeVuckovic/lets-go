package server

import (
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/internal/db"
	mw "github.com/DjordjeVuckovic/lets-go/pkg/api/middleware"
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
	Echo *echo.Echo
	cfg  *Config
}

func NewServer(e *echo.Echo, ec *Config, database *db.Database) *Server {
	e.DisableHTTP2 = !ec.UseHttp2

	s := &Server{
		Echo: e,
		cfg:  ec,
	}

	s.setupMiddlewares()
	s.setupHealthChecks(database)

	return s
}

func (s *Server) setupMiddlewares() {
	s.Echo.Use(mw.Logger())
	s.Echo.Use(middleware.Recover())
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: s.cfg.CorsOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}

func (s *Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := s.Echo.Start(":" + s.cfg.Port); err != nil && err != http.ErrServerClosed {
			s.Echo.Logger.Fatal("shutting down the server")
		}
	}()
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), GracefulShutdownTimeout)
	defer cancel()
	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
		return err
	}
	return nil
}

func (s *Server) setupHealthChecks(db *db.Database) {
	s.Echo.GET("/health", func(c echo.Context) error {
		ctx := c.Request().Context()
		err := db.PingCtx(ctx)
		if err != nil {
			return err
		}
		return c.JSON(200, "ok")
	})
	s.Echo.GET("/ready", func(c echo.Context) error {
		ctx := c.Request().Context()
		err := db.PingCtx(ctx)
		if err != nil {
			return err
		}
		return c.JSON(200, "ok")
	})
}
