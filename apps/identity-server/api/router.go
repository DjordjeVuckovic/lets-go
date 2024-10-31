package api

import (
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/internal/db"
	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	e  *echo.Echo
	db *db.Database
}

func NewAuthRouter(e *echo.Echo, db *db.Database) *AuthRouter {
	a := &AuthRouter{e: e, db: db}
	return a
}

func (ar *AuthRouter) RegisterRoutes() {
	g := ar.e.Group("api/v1/auth")
	g.POST("/register", ar.handleRegister)
	g.POST("/login", ar.handleLogin)
}

func (ar *AuthRouter) handleLogin(c echo.Context) error {
	return c.JSON(200, "Hello, World!")
}

func (ar *AuthRouter) handleRegister(c echo.Context) error {
	ctx := c.Request().Context()
	RegisterUser(ctx, ar.db)
	return c.JSON(200, "Hello, World!")
}
