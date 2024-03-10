package server

import (
	inventoryHttp "firdausyusofs/kem_digital/internal/inventory/http"
	inventoryRepo "firdausyusofs/kem_digital/internal/inventory/repository"
	inventoryUC "firdausyusofs/kem_digital/internal/inventory/usecase"

	"github.com/labstack/echo/v4"
)

func (s *Server) MakeHandlers(e *echo.Echo) error {
	// Init repositories
	inventoryRepo := inventoryRepo.NewInventoryRepository(s.db)

	// Init usecases
	inventoryUC := inventoryUC.NewInventoryUseCase(s.cfg, inventoryRepo, s.logger)

	// Init handlers
	inventoryHandlers := inventoryHttp.NewInventoryHandler(s.cfg, inventoryUC, s.logger)

	api := e.Group("/api")

	inventoryHttp.MakeInventoryRoutes(api, inventoryHandlers)

	return nil
}
