package http

import (
	"firdausyusofs/kem_digital/internal/supplier"

	"github.com/labstack/echo/v4"
)

func MakeSupplierRoutes(group *echo.Group, h supplier.Handlers) {
	group.GET("/suppliers", h.GetSuppliers())
}
