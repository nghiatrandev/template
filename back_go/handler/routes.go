package handler

import "github.com/labstack/echo/v4"

func (s *Server) buildRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/register", s.Register)

	return e
}
