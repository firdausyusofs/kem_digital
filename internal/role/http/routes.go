package http

import (
	"firdausyusofs/kem_digital/internal/role"

	"github.com/labstack/echo/v4"
)

func MakeRoleRoutes(group *echo.Group, h role.Handlers) {
	group.GET("", h.GetRoles())
}
