package supplier

import "github.com/labstack/echo/v4"

type Handlers interface {
	GetSuppliers() echo.HandlerFunc
}
