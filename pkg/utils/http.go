package utils

import "github.com/labstack/echo/v4"

func ReadBody(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}

	return ValidateStruct(c.Request().Context(), req)
}
