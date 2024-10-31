package middleware

import (
	"context"
	"github.com/DjordjeVuckovic/lets-go/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
)

type LoggerOpts func(*middleware.RequestLoggerConfig)

func Logger(opts ...LoggerOpts) echo.MiddlewareFunc {
	o := defaultOpt()
	for _, opt := range opts {
		opt(&o)
	}

	return middleware.RequestLoggerWithConfig(o)
}

func defaultOpt() middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogLatency:  true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.GetLogger().LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					logger.String("uri", v.URI),
					logger.Int("status", v.Status),
				)
			} else {
				logger.GetLogger().LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					logger.String("uri", v.URI),
					logger.Int("status", v.Status),
					logger.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}
}