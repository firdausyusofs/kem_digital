package http

import (
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/role"
	"firdausyusofs/kem_digital/pkg/api_response"
	"firdausyusofs/kem_digital/pkg/logger"

	"github.com/labstack/echo/v4"
)

type roleHandler struct {
	cfg    *config.Config
	roleUC role.UseCase
	logger logger.Logger
}

func NewRoleHandler(cfg *config.Config, roleUC role.UseCase, logger logger.Logger) role.Handlers {
	return &roleHandler{
		cfg:    cfg,
		roleUC: roleUC,
		logger: logger,
	}
}

func (h *roleHandler) GetRoles() echo.HandlerFunc {
	return func(c echo.Context) error {
		roles, err := h.roleUC.GetRoles(c.Request().Context())
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(404, "Failed to retrieve roles", &errMsg))
		}

		return c.JSON(api_response.MakeSuccessResponse(200, "Successfully retrieved roles", &roles))
	}
}
