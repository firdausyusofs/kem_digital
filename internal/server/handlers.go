package server

import "github.com/labstack/echo/v4"

func (s *Server) MakeHandlers(e *echo.Echo) error {
	hello := e.Group("/hello")

	hello.GET("", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return nil
}
