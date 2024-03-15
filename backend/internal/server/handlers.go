package server

import (
	authHttp "firdausyusofs/kem_digital/internal/auth/http"
	authRepo "firdausyusofs/kem_digital/internal/auth/repository"
	authUC "firdausyusofs/kem_digital/internal/auth/usecase"
	inventoryHttp "firdausyusofs/kem_digital/internal/inventory/http"
	inventoryRepo "firdausyusofs/kem_digital/internal/inventory/repository"
	inventoryUC "firdausyusofs/kem_digital/internal/inventory/usecase"
	roleHttp "firdausyusofs/kem_digital/internal/role/http"
	roleRepo "firdausyusofs/kem_digital/internal/role/repository"
	roleUC "firdausyusofs/kem_digital/internal/role/usecase"
	supplierHttp "firdausyusofs/kem_digital/internal/supplier/http"
	supplierRepo "firdausyusofs/kem_digital/internal/supplier/repository"
	supplierUC "firdausyusofs/kem_digital/internal/supplier/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MakeHandlers(e *echo.Echo) error {
	// Init repositories
	inventoryRepo := inventoryRepo.NewInventoryRepository(s.db)
	supplierRepo := supplierRepo.NewSupplierRepository(s.db)
	authRepo := authRepo.NewAuthRepository(s.db)
	roleRepo := roleRepo.NewRoleRepository(s.db)

	// Init usecases
	inventoryUC := inventoryUC.NewInventoryUseCase(s.cfg, inventoryRepo, s.logger)
	supplierUC := supplierUC.NewSupplierUseCase(s.cfg, supplierRepo, s.logger)
	authUC := authUC.NewAuthUseCase(s.cfg, authRepo, s.logger)
	roleUC := roleUC.NewRoleUseCase(s.cfg, roleRepo, s.logger)

	// Init handlers
	inventoryHandlers := inventoryHttp.NewInventoryHandler(s.cfg, inventoryUC, s.logger)
	supplierHandlers := supplierHttp.NewSupplierHandler(s.cfg, supplierUC, s.logger)
	authHandlers := authHttp.NewAuthHandler(s.cfg, authUC, s.logger)
	roleHandlers := roleHttp.NewRoleHandler(s.cfg, roleUC, s.logger)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))

	api := e.Group("/api")

	inventoryHttp.MakeInventoryRoutes(api, inventoryHandlers)
	supplierHttp.MakeSupplierRoutes(api, supplierHandlers)

	authGroup := api.Group("/auth")
	authHttp.MakeAuthRoutes(authGroup, authHandlers)

	roleGroup := api.Group("/roles")
	roleHttp.MakeRoleRoutes(roleGroup, roleHandlers)

	return nil
}
