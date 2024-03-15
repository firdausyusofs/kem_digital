package http

import (
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/supplier"
	"firdausyusofs/kem_digital/pkg/api_response"
	"firdausyusofs/kem_digital/pkg/logger"

	"github.com/labstack/echo/v4"
)

type supplierHandler struct {
	cfg        *config.Config
	supplierUC supplier.UseCase
	logger     logger.Logger
}

func NewSupplierHandler(cfg *config.Config, supplierUC supplier.UseCase, logger logger.Logger) supplier.Handlers {
	return &supplierHandler{
		cfg:        cfg,
		supplierUC: supplierUC,
		logger:     logger,
	}
}

func (h *supplierHandler) GetSuppliers() echo.HandlerFunc {
	return func(c echo.Context) error {
		suppliers, err := h.supplierUC.GetSuppliers(c.Request().Context())
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(404, "Failed to retrieve suppliers", &errMsg))
		}

		return c.JSON(api_response.MakeSuccessResponse(200, "Successfully retrieved suppliers", &suppliers))
	}
}
