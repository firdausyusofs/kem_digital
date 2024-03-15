package role

import "github.com/labstack/echo/v4"

type Handlers interface {
	GetRoles() echo.HandlerFunc
}
