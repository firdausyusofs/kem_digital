package http

import (
	"firdausyusofs/kem_digital/internal/inventory"

	"github.com/labstack/echo/v4"
)

func MakeInventoryRoutes(group *echo.Group, h inventory.Handlers) {
	group.GET("/inventory", h.GetProducts())
	group.GET("/inventory/:product_id", h.GetProductByID())
	group.POST("/add-inventory", h.CreateProduct())
	group.PUT("/update-inventory/:product_id", h.UpdateProduct())
	group.DELETE("/delete-inventory/:product_id", h.DeleteProduct())
}
