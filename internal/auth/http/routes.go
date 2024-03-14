package http

import (
	"firdausyusofs/kem_digital/internal/auth"

	"github.com/labstack/echo/v4"
)

func MakeAuthRoutes(group *echo.Group, h auth.Handlers) {
	group.POST("/register", h.Register())
	group.POST("/login", h.Login())
}
