package inventory

import "github.com/labstack/echo/v4"

type Handlers interface {
	GetProducts() echo.HandlerFunc
	GetProductByID() echo.HandlerFunc
	CreateProduct() echo.HandlerFunc
}
