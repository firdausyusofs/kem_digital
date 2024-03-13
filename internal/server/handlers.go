package server

import (
	inventoryHttp "firdausyusofs/kem_digital/internal/inventory/http"
	inventoryRepo "firdausyusofs/kem_digital/internal/inventory/repository"
	inventoryUC "firdausyusofs/kem_digital/internal/inventory/usecase"
	supplierHttp "firdausyusofs/kem_digital/internal/supplier/http"
	supplierRepo "firdausyusofs/kem_digital/internal/supplier/repository"
	supplierUC "firdausyusofs/kem_digital/internal/supplier/usecase"

	"github.com/labstack/echo/v4"
)

func (s *Server) MakeHandlers(e *echo.Echo) error {
	// Init repositories
	inventoryRepo := inventoryRepo.NewInventoryRepository(s.db)
	supplierRepo := supplierRepo.NewSupplierRepository(s.db)

	// Init usecases
	inventoryUC := inventoryUC.NewInventoryUseCase(s.cfg, inventoryRepo, s.logger)
	supplierUC := supplierUC.NewSupplierUseCase(s.cfg, supplierRepo, s.logger)

	// Init handlers
	inventoryHandlers := inventoryHttp.NewInventoryHandler(s.cfg, inventoryUC, s.logger)
	supplierHandlers := supplierHttp.NewSupplierHandler(s.cfg, supplierUC, s.logger)

	api := e.Group("/api")

	inventoryHttp.MakeInventoryRoutes(api, inventoryHandlers)
	supplierHttp.MakeSupplierRoutes(api, supplierHandlers)

	return nil
}
